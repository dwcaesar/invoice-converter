package main

import (
	"daquam/metric"
	"daquam/report"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Metric string

const (
	Completeness Metric = "Completeness"
	Consistency  Metric = "Consistency"
)

type MetricConfig struct {
	Id     Metric   `yaml:"id"`
	Fields []string `yaml:"fields"`
}

type Item struct {
	Name      string `xml:"name"`
	Amount    string `xml:"amount"`
	ItemPrice string `xml:"itemPrice"`
	Vat       string `xml:"vat"`
}

type Address struct {
	Name   string `xml:"name"`
	Street string `xml:"street"`
	Zip    string `xml:"zip"`
	City   string `xml:"city"`
}

type Invoice struct {
	XMLName         xml.Name `xml:"invoice"`
	InvoiceNumber   string   `xml:"invoiceNumber"`
	BillingAddress  Address  `xml:"billingAddress"`
	ShippingAddress Address  `xml:"shippingAddress"`
	PaymentMethod   string   `xml:"paymentMethod"`
	Items           []Item   `xml:"items>item"`
	Netto           string   `xml:"netto"`
	Brutto          string   `xml:"brutto"`
}

// Completeness is given when all value are set and not empty
func validateInvoiceCompleteness(fields []string, record map[string]interface{}, csvReport *report.Report) bool {

	for _, field := range fields {
		value, exists := record[field]
		csvReport.NewEntry(field, "FieldExists")
		if !exists || value == nil {
			log.Default().Printf("the record's field %s was incomplete", field)
			csvReport.CloseEntry(nil, false, "field is missing")
			return false
		} else {
			// instead of the actual value, an empty string
			// otherwise, complex objects might be serialized in the report
			csvReport.CloseEntry("", true, "")
		}

		switch field {
		case "BillingAddress", "ShippingAddress":
			if !metric.IsAddressComplete(value, field, csvReport) {
				log.Default().Printf("the record's field %s was incomplete", field)
				return false
			}
		case "Items":
			if !metric.IsItemsComplete(value, field, csvReport) {
				log.Default().Printf("the record's field %s was incomplete", field)
				return false
			}
		default:
			csvReport.NewEntry(field, "FieldHasValue")
			str := strings.TrimSpace(fmt.Sprintf("%s", value))
			if len(str) == 0 {
				log.Default().Printf("the record's field %s was incomplete", field)
				csvReport.CloseEntry(str, false, "")
				return false
			} else {
				csvReport.CloseEntry(str, true, "")
			}
		}
	}
	return true
}

// We consider Intra-record and Format Consistencies here
func validateInvoiceConsistency(fields []string, record map[string]interface{}, csvReport *report.Report) bool {
	for _, field := range fields {
		switch field {
		case "Brutto", "Netto":
			if !metric.IsBruttoNettoConsistent(record, csvReport) {
				return false
			}
		case "Items":
			if !metric.IsNettoPriceConsistent(record, csvReport) {
				return false
			}
		default:
			return false
		}
	}
	return true
}

func getCalcFunctionForMetric(metric Metric) (func(fields []string, record map[string]interface{}, csvReport *report.Report) bool, error) {
	switch metric {
	case Completeness:
		return validateInvoiceCompleteness, nil
	case Consistency:
		return validateInvoiceConsistency, nil
	default:
		return nil, errors.New("unexected metric")
	}
}

func calculateMetric(metric Metric, fields []string, records []map[string]interface{}, csvReport *report.Report) float64 {
	totalRecords := len(records)
	if totalRecords == 0 {
		log.Println("no records found")
		return 0
	}

	validRecords := 0

	validationFunc, err := getCalcFunctionForMetric(metric)
	if err != nil {
		log.Printf("coudnt determine the validation function for %s due to %s", metric, err)
		return 0
	}
	for _, rec := range records {
		if validationFunc(fields, rec, csvReport) {
			validRecords++
		}
	}

	log.Printf("calculate %s by considering fields %s : records valid %d, records total %d\n", metric, fields, validRecords, totalRecords)

	return float64(validRecords) / float64(totalRecords)
}

func convertInvoiceToMap(invoice Invoice) (map[string]interface{}, error) {
	var result map[string]interface{}
	jsonBytes, err := json.Marshal(invoice)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBytes, &result)
	return result, err
}

func main() {
	files, err := filepath.Glob("assets/invoice*.xml")
	if err != nil {
		log.Fatalf("Error finding XML files: %v", err) // Fatalf logs the error and exits
	}

	if len(files) == 0 {
		log.Println("No XML files found in the current folder.")
		return
	}

	var invoices []Invoice

	// Read and parse each XML file
	for _, file := range files {
		log.Println("Processing file:", file)
		if inv, err := parseInvoiceXML(file); err == nil {
			invoices = append(invoices, inv)
		}
	}

	var invoiceRecords []map[string]interface{}

	for _, i := range invoices {
		if result, err := convertInvoiceToMap(i); err == nil {
			invoiceRecords = append(invoiceRecords, result)
		}
	}

	config, err := parseYamlConfig("assets/config.yaml")
	if err != nil {
		log.Default().Fatal("Error finding yaml file:", err)
	}

	csvReport, err := report.NewReport("assets")
	defer func(csvReport *report.Report) {
		err := csvReport.CloseReport()
		if err != nil {
			log.Default().Printf("Error closing report: %v", err)
		}
	}(csvReport)

	for _, conf := range config {
		metricValue := calculateMetric(conf.Id, conf.Fields, invoiceRecords, csvReport) * 100
		log.Printf("Metric %s: %.2f%%\n", conf.Id, metricValue)
	}
}

func parseYamlConfig(filename string) ([]MetricConfig, error) {
	//Read yaml
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read file %s: %v\n", filename, err)
		return nil, err
	}

	// Parse the XML into the Employees struct
	var metricsConfig []MetricConfig
	err = yaml.Unmarshal(data, &metricsConfig)
	if err != nil {
		log.Printf("Failed to parse XML in %s: %v\n", filename, err)
		return nil, err
	}

	return metricsConfig, nil
}

func parseInvoiceXML(filename string) (Invoice, error) {
	// Read the XML file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read file %s: %v\n", filename, err)
		return Invoice{}, err
	}

	// Parse the XML into the Employees struct
	var invoice Invoice
	err = xml.Unmarshal(data, &invoice)
	if err != nil {
		log.Printf("Failed to parse XML in %s: %v\n", filename, err)
		return Invoice{}, err
	}

	log.Printf("read invoice %s", invoice)
	return invoice, nil
}
