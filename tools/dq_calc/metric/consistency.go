package metric

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

func isPriceValueConsistent(price string) bool {
	if _, err := strconv.ParseFloat(price, 64); err == nil {
		var spiltByDec = strings.Split(price, ".")
		length := len(spiltByDec[1])
		return len(spiltByDec) == 2 && length >= 2
	}

	return false
}

func IsNettoPriceConsistent(record map[string]interface{}) bool {
	nettoValue, exists := record["Netto"]
	if !exists {
		return false
	}
	nettoStringValue, canBeConverted := nettoValue.(string)
	if !canBeConverted || !isPriceValueConsistent(nettoStringValue) {
		return false
	}

	nettoPrice, err := strconv.ParseFloat(nettoStringValue, 64)
	if err != nil {
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
			return false
		}

		itemPriceString, isConverted := itemPrice.(string)
		if !isConverted {
			return false
		}

		itemPriceValue, err := strconv.ParseFloat(itemPriceString, 64)
		if err != nil {
			return false
		}

		itemAmount, exists := itemValue["Amount"]
		if !exists {
			return false
		}

		itemAmountString, isConverted := itemAmount.(string)
		if !isConverted {
			return false
		}

		itemAmountValue, err := strconv.ParseUint(itemAmountString, 10, 32)
		if err != nil {
			return false
		}
		calculatedNettoPrice += float64(itemAmountValue) * itemPriceValue

	}
	return nettoPrice == calculatedNettoPrice
}

func IsBruttoNettoConsistent(record map[string]interface{}) bool {
	nettoValue, exists := record["Netto"]
	if !exists {
		return false
	}
	nettoStringValue, canBeConverted := nettoValue.(string)
	if !canBeConverted || !isPriceValueConsistent(nettoStringValue) {
		return false
	}

	nettoPrice, err := strconv.ParseFloat(nettoStringValue, 64)
	if err != nil {
		return false
	}

	bruttoValue, exists := record["Brutto"]
	if !exists {
		return false
	}
	bruttoStringValue, canBeConverted := bruttoValue.(string)
	if !canBeConverted || !isPriceValueConsistent(bruttoStringValue) {
		return false
	}

	bruttoPrice, err := strconv.ParseFloat(bruttoStringValue, 64)
	if err != nil {
		return false
	}
	return bruttoPrice > nettoPrice
}
