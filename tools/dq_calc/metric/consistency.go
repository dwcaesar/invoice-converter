package metric

import (
	"daquam/report"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func isPriceValueConsistent(price string, csvReport *report.Report) bool {
	if _, err := strconv.ParseFloat(price, 64); err == nil {
		var spiltByDec = strings.Split(price, ".")
		length := len(spiltByDec[1])
		hasTwoDecimalPlaces := len(spiltByDec) == 2 && length >= 2
		if !hasTwoDecimalPlaces {
			csvReport.CloseEntry(price, false, "should have two decimal places")
		}
		return hasTwoDecimalPlaces
	}
	csvReport.CloseEntry(price, false, "did not contain a valid price")
	return false
}

func IsNettoPriceConsistent(record map[string]interface{}, csvReport *report.Report) bool {
	nettoValue, exists := record["Netto"]
	csvReport.NewEntry("Netto", "IsNettoPriceConsistent")
	if !exists {
		csvReport.CloseEntry(nil, false, "field is missing")
		return false
	}
	nettoStringValue, canBeConverted := nettoValue.(string)
	if !canBeConverted {
		csvReport.CloseEntry(nettoValue, false, "field is not a string")
		return false
	}
	if !isPriceValueConsistent(nettoStringValue, csvReport) {
		return false
	}

	nettoPrice, err := strconv.ParseFloat(nettoStringValue, 64)
	if err != nil {
		csvReport.CloseEntry(nettoStringValue, false, "field is not a float64")
		return false
	}

	items, isConverted := record["Items"].([]interface{})
	if !isConverted {
		log.Panicf("cannot convert to items[] %s", reflect.TypeOf(record["Items"]))
		return false
	}

	calculatedNettoPrice := 0.0

	for _, item := range items {
		itemValue, isConverted := item.(map[string]interface{})
		if !isConverted {
			log.Panicf("cannot convert to items[] %s", reflect.TypeOf(item))
			return false
		}

		itemPrice, exists := itemValue["ItemPrice"]
		if !exists {
			csvReport.CloseEntry(nettoStringValue, false, "invalid item found")
			return false
		}

		itemPriceString, isConverted := itemPrice.(string)
		if !isConverted {
			csvReport.CloseEntry(nettoStringValue, false, "item without price found")
			return false
		}

		itemPriceValue, err := strconv.ParseFloat(itemPriceString, 64)
		if err != nil {
			csvReport.CloseEntry(nettoStringValue, false, "item with invalid price found")
			return false
		}

		itemAmount, exists := itemValue["Amount"]
		if !exists {
			csvReport.CloseEntry(nettoStringValue, false, "item without amount found")
			return false
		}

		itemAmountString, isConverted := itemAmount.(string)
		if !isConverted {
			csvReport.CloseEntry(nettoStringValue, false, "item with invalid amount found")
			return false
		}

		itemAmountValue, err := strconv.ParseUint(itemAmountString, 10, 32)
		if err != nil {
			csvReport.CloseEntry(nettoStringValue, false, "item with invalid amount found")
			return false
		}
		calculatedNettoPrice += float64(itemAmountValue) * itemPriceValue

	}
	pricesMatch := nettoPrice == calculatedNettoPrice
	csvReport.CloseEntry(nettoStringValue, pricesMatch, fmt.Sprintf("expected total netto: %f", calculatedNettoPrice))
	return pricesMatch
}

func IsBruttoNettoConsistent(record map[string]interface{}, report *report.Report) bool {
	report.NewEntry("Netto", "IsBruttoNettoConsistent")
	nettoValue, exists := record["Netto"]
	if !exists {
		report.CloseEntry(nil, false, "field is missing")
		return false
	}
	nettoStringValue, canBeConverted := nettoValue.(string)
	if !canBeConverted {
		report.CloseEntry(nettoValue, false, "value is invalid")
		return false
	}
	if !isPriceValueConsistent(nettoStringValue, report) {
		return false
	}

	nettoPrice, err := strconv.ParseFloat(nettoStringValue, 64)
	if err != nil {
		report.CloseEntry(nettoStringValue, false, "value is invalid")
		return false
	}

	bruttoValue, exists := record["Brutto"]
	if !exists {
		report.CloseEntry(nil, false, "brutto is missing")
		return false
	}
	bruttoStringValue, canBeConverted := bruttoValue.(string)
	if !canBeConverted {
		report.CloseEntry(bruttoValue, false, "value is invalid")
		return false
	}
	if !isPriceValueConsistent(bruttoStringValue, report) {
		return false
	}

	bruttoPrice, err := strconv.ParseFloat(bruttoStringValue, 64)
	if err != nil {
		report.CloseEntry(bruttoStringValue, false, "value is invalid")
		return false
	}
	isBruttoGreaterNetto := bruttoPrice > nettoPrice
	report.CloseEntry(bruttoStringValue, isBruttoGreaterNetto, fmt.Sprintf("nettoPrice: %f", nettoPrice))
	return isBruttoGreaterNetto
}
