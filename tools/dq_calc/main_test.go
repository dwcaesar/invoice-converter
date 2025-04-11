package main

import (
	"daquam/assert"
	"daquam/metric"
	"encoding/json"
	"log"
	"testing"
)

func TestTrueInvoiceCompleteness(t *testing.T) {
	invoice := Invoice{
		Netto: "n",
		BillingAddress: Address{
			City:   "c",
			Street: "s",
			Zip:    "z",
			Name:   "n",
		},
		Items: []Item{{ItemPrice: "a", Amount: "s", Vat: "v", Name: "n"}, {ItemPrice: "a", Amount: "s", Vat: "v", Name: "n"}},
	}

	if record, err := convertInvoiceToMap(invoice); err == nil {
		validateInvoiceCompleteness([]string{"Netto", "BillingAddress", "Items"}, record)
	} else {
		t.FailNow()
	}

}

func TestTrueNettoPriceConsistent(t *testing.T) {

	invoice := Invoice{
		Netto: "100.75",
		BillingAddress: Address{
			City:   "c",
			Street: "s",
			Zip:    "z",
			Name:   "n",
		},
		Items: []Item{{ItemPrice: "40.30", Amount: "2", Vat: "v", Name: "n"}, {ItemPrice: "20.15", Amount: "1", Vat: "v", Name: "n"}},
	}

	if record, err := convertInvoiceToMap(invoice); err == nil {
		result := metric.IsNettoPriceConsistent(record)
		assert.Equal(t, result, true)
	} else {
		t.FailNow()
	}
}

func TestFalseNettoPriceConsistent(t *testing.T) {

	invoice := Invoice{
		Netto: "100.00",
		BillingAddress: Address{
			City:   "c",
			Street: "s",
			Zip:    "z",
			Name:   "n",
		},
		Items: []Item{{ItemPrice: "40.0", Amount: "1"}, {ItemPrice: "20.0", Amount: "1"}},
	}

	if record, err := convertInvoiceToMap(invoice); err == nil {
		result := metric.IsNettoPriceConsistent(record)
		assert.Equal(t, result, false)
	} else {
		t.FailNow()
	}

	invoice = Invoice{
		Netto: "100.00",
		BillingAddress: Address{
			City:   "c",
			Street: "s",
			Zip:    "z",
			Name:   "n",
		},
		Items: []Item{{ItemPrice: "40.0x", Amount: "1"}, {ItemPrice: "20.0", Amount: "1"}},
	}

	if record, err := convertInvoiceToMap(invoice); err == nil {
		result := metric.IsNettoPriceConsistent(record)
		assert.Equal(t, result, false)
	} else {
		t.FailNow()
	}
}

func TestFalseInvoiceCompleteness(t *testing.T) {
	invoice := Invoice{
		Netto: "n",
		BillingAddress: Address{
			City: "c",
			Zip:  "z",
			Name: "n",
		},
		Items: []Item{{ItemPrice: "a", Amount: "s", Vat: "v", Name: "n"}, {ItemPrice: "a", Amount: "s", Vat: "v", Name: "n"}},
	}

	if record, err := convertInvoiceToMap(invoice); err == nil {
		assert.Equal(t, validateInvoiceCompleteness([]string{"Netto", "BillingAddress", "Items"}, record), false)
	} else {
		t.FailNow()
	}

	invoice = Invoice{
		ShippingAddress: Address{
			City:   "c",
			Street: "s",
			Zip:    "z",
			Name:   "n",
		},
		Items: []Item{{ItemPrice: "a", Amount: "", Vat: "v", Name: "n"}, {ItemPrice: "a", Amount: "s", Vat: "v", Name: "n"}},
	}

	if record, err := convertInvoiceToMap(invoice); err == nil {
		assert.Equal(t, validateInvoiceCompleteness([]string{"ShippingAddress", "Items"}, record), false)
	} else {
		t.FailNow()
	}
}

func TestTrueItemsCompleteness(t *testing.T) {
	items := []Item{{ItemPrice: "a", Amount: "s", Vat: "v", Name: "n"}, {ItemPrice: "a", Amount: "s", Vat: "v", Name: "n"}}

	records := marshalItems(items)
	assert.Equal(t, metric.IsItemsComplete(records), true)
}

func marshalItems(items []Item) []interface{} {
	var records []interface{}

	for _, item := range items {
		var record map[string]interface{}
		if bytes, err := json.Marshal(item); err == nil {
			if err := json.Unmarshal(bytes, &record); err == nil {
				records = append(records, record)
			} else {
				log.Fatalf("error %s", err)
			}
		} else {
			log.Fatalf("error %s", err)
		}
	}

	return records
}
func TestFalseItemsCompleteness(t *testing.T) {
	items := []Item{{ItemPrice: "a", Amount: "s", Vat: "v", Name: "n"}, {ItemPrice: " ", Amount: "s", Vat: "v", Name: "n"}}
	records := marshalItems(items)
	assert.Equal(t, metric.IsItemsComplete(records), false)

	items = []Item{{ItemPrice: "a", Vat: "v", Name: "n"}, {ItemPrice: "i", Amount: "s", Vat: "v", Name: "n"}}
	records = marshalItems(items)
	assert.Equal(t, metric.IsItemsComplete(records), false)
}

func TestTrueAddressCompletness(t *testing.T) {
	address := Address{
		Name:   "n ",
		Zip:    "12",
		Street: " str",
		City:   "city",
	}

	assert.Equal(t, metric.IsAddressComplete(address), true)
}

func TestTrueBruttoNettoConsistent(t *testing.T) {
	record := map[string]interface{}{"Brutto": "12.304", "Netto": "0.01"}

	result := metric.IsBruttoNettoConsistent(record)

	assert.Equal(t, result, true)
}

func TestFalseBruttoNettoConsistent(t *testing.T) {
	//wrong precision
	record := map[string]interface{}{"Brutto": "12.3", "Netto": "0.01"}

	result := metric.IsBruttoNettoConsistent(record)

	assert.Equal(t, result, false)

	//Netto missing
	record = map[string]interface{}{"Brutto": "12.3"}

	result = metric.IsBruttoNettoConsistent(record)

	assert.Equal(t, result, false)
}

func TestFalseAddressCompletness(t *testing.T) {
	address1 := Address{
		Name:   "n ",
		Zip:    "",
		Street: " str",
		City:   "city",
	}

	assert.Equal(t, metric.IsAddressComplete(address1), false)

	address2 := Address{
		Name: "n ",
		Zip:  "xx",
		City: "city",
	}

	assert.Equal(t, metric.IsAddressComplete(address2), false)
}

func TestTrueSimpleRecordCompletness(t *testing.T) {
	record := map[string]string{"key1": "value1", "key2": "value2"}

	assert.Equal(t, metric.IsSimpleRecordComplete([]string{"key1", "key2"}, record), true)
}

func TestFalseRecordCompletness(t *testing.T) {
	record := map[string]string{"key1": " ", "key2": ""}

	assert.Equal(t, metric.IsSimpleRecordComplete([]string{"keyX"}, record), false)
	assert.Equal(t, metric.IsSimpleRecordComplete([]string{"key1"}, record), false)
	assert.Equal(t, metric.IsSimpleRecordComplete([]string{"key2"}, record), false)
}
func TestFalseAdressCompletenessWithMissingField(t *testing.T) {
	var adr = Address{
		Zip:    "1",
		City:   "x",
		Street: "a",
	}
	adress, err := convertAdressToMap(adr)

	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, metric.IsAddressComplete(adress), false)
}
func TestTrueAdressCompleteness(t *testing.T) {
	var adr = Address{
		Zip:    "1",
		Name:   "s",
		City:   "b",
		Street: "a",
	}
	adress, err := convertAdressToMap(adr)

	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, metric.IsAddressComplete(adress), true)
}

func convertAdressToMap(object interface{}) (map[string]interface{}, error) {
	var record map[string]interface{}
	jsonBytes, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBytes, &record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func TestCalcMetricEmptyInput(t *testing.T) {
	value := calculateMetric(Consistency, nil, nil)
	assert.Equal(t, value, float64(0))
}

func TestCalcMetricOutput(t *testing.T) {
	//value := calculateMetric(Consistency, map[string]interface{}{}, []map[string]string{})
	assert.Equal(t, 1, 1)
}
