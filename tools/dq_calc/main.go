package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	layoutISO    = "2006-01-02"
	emailPattern = `^[\w\.-]+@[\w\.-]+\.\w+$`
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

// Employee struct to hold sample data
type Employee struct {
	Id        string `xml:"id"`
	Birthdate string `xml:"birthdate"`
	Hiredate  string `xml:"hiredate"`
	Email     string `xml:"email"`
}

type EmployeesXML struct {
	XMLName   xml.Name   `xml:"employees"`
	Employees []Employee `xml:"employee"`
}

// checks if the email follows a given pattern
func validateEmail(email string) bool {
	re := regexp.MustCompile(emailPattern)
	return re.MatchString(email)
}

// Completeness is given when all value are set and not empty
func validateCompleteness(fields []string, record map[string]string) bool {
	for _, field := range fields {
		value, exists := record[field]
		if !exists || len(strings.TrimSpace(value)) == 0 {
			return false
		}
	}
	return true
}

// We considfer Intra-record and Format Consistenies here
func validateConsistency(fields []string, record map[string]string) bool {

	for _, field := range fields {
		value, exists := record[field]

		if !exists {
			return false
		}

		switch field {
		case "Email":
			if !validateEmail(value) {
				return false
			}
		case "Birthdate":
			if !validateHiredateBirthdateConsistency(record) {
				return false
			}
		case "Hiredate":
			if !validateHiredateBirthdateConsistency(record) {
				return false
			}
		}
	}
	return true
}

func validateHiredateBirthdateConsistency(record map[string]string) bool {
	valueHire, existsHire := record["Hiredate"]
	valueBirth, existsBirth := record["Birthdate"]

	if !existsBirth || !existsHire {
		return false
	}

	birth, err1 := time.Parse(layoutISO, valueBirth)
	hire, err2 := time.Parse(layoutISO, valueHire)

	return err1 == nil && err2 == nil && birth.Before(hire)
}

func getCalcFunctionForMetric(metric Metric) (func(fields []string, record map[string]string) bool, error) {
	switch metric {
	case Completeness:
		return validateCompleteness, nil
	case Consistency:
		return validateConsistency, nil
	default:
		return nil, errors.New("unexected metric")
	}
}

func calculateMetric(metric Metric, fields []string, records []map[string]string) float64 {
	totalRecords := len(records)
	if totalRecords == 0 {
		log.Println("empty employees array")
		return 0
	}

	validRecords := 0

	vlaidationFunc, err := getCalcFunctionForMetric(metric)
	if err != nil {
		log.Printf("coudnt determine the validation function for %s due to %s", metric, err)
		return 0
	}
	for _, rec := range records {
		if vlaidationFunc(fields, rec) {
			validRecords++
		}
	}

	log.Printf("calculate %s by considering fields %s : records valid %d, records total %d\n", metric, fields, validRecords, totalRecords)

	return float64(validRecords) / float64(totalRecords)
}

// Metric value calculated as metric = validRecords / totalRecords.
// Determine the amount of valid records by applying validationRecordFunc to every record in []Employee
func convertEmployeeToMap(emp Employee) (map[string]string, error) {
	var result map[string]string
	jsonBytes, err := json.Marshal(emp)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	files, err := filepath.Glob("assets/*.xml")
	if err != nil {
		log.Println("Error finding XML files:", err)
		return
	} else if len(files) == 0 {
		log.Println("No XML files found in the current folder.")
		return
	}

	var employees []Employee

	// Read and parse each XML file
	for _, file := range files {
		log.Println("Processing file:", file)
		if empl, err := parseXML(file); err == nil {
			employees = append(employees, empl...)
		}
	}

	var employeesMap []map[string]string

	for _, emp := range employees {
		if result, err := convertEmployeeToMap(emp); err == nil {
			employeesMap = append(employeesMap, result)
		}
	}

	config, err := parseYamlConfig("assets/config.yaml")
	if err != nil {
		log.Default().Fatal("Error finding yaml file:", err)
	}

	for _, conf := range config {
		metricValue := calculateMetric(conf.Id, conf.Fields, employeesMap) * 100
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

func parseXML(filename string) ([]Employee, error) {
	// Read the XML file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read file %s: %v\n", filename, err)
		return nil, err
	}

	// Parse the XML into the Employees struct
	var employees EmployeesXML
	err = xml.Unmarshal(data, &employees)
	if err != nil {
		log.Printf("Failed to parse XML in %s: %v\n", filename, err)
		return nil, err
	}

	log.Printf("parsed %d records", len(employees.Employees))
	return employees.Employees, nil
}
