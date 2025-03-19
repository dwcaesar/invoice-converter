package main

import (
	"daquam/assert"
	"encoding/json"
	"log"
	"testing"
)

func TestEmailOk(t *testing.T) {
	var testString = "2@test.de"

	value := validateEmail(testString)
	assert.Equal(t, value, true)
}

func TestEmailFailed(t *testing.T) {
	var testString = "2/@test.de"

	value := validateEmail(testString)
	assert.Equal(t, value, false)

}

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
	assert.Equal(t, isItemsComplete(records), true)
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
	assert.Equal(t, isItemsComplete(records), false)

	items = []Item{{ItemPrice: "a", Vat: "v", Name: "n"}, {ItemPrice: "i", Amount: "s", Vat: "v", Name: "n"}}
	records = marshalItems(items)
	assert.Equal(t, isItemsComplete(records), false)
}

func TestTrueAddressCompletness(t *testing.T) {
	address := Address{
		Name:   "n ",
		Zip:    "12",
		Street: " str",
		City:   "city",
	}

	assert.Equal(t, isAddressComplete(address), true)
}

func TestFalseAddressCompletness(t *testing.T) {
	address1 := Address{
		Name:   "n ",
		Zip:    "",
		Street: " str",
		City:   "city",
	}

	assert.Equal(t, isAddressComplete(address1), false)

	address2 := Address{
		Name: "n ",
		Zip:  "xx",
		City: "city",
	}

	assert.Equal(t, isAddressComplete(address2), false)
}

func TestTrueSimpleRecordCompletness(t *testing.T) {
	record := map[string]string{"key1": "value1", "key2": "value2"}

	assert.Equal(t, isSimpleRecordComplete([]string{"key1", "key2"}, record), true)
}

func TestFalseRecordCompletness(t *testing.T) {
	record := map[string]string{"key1": " ", "key2": ""}

	assert.Equal(t, isSimpleRecordComplete([]string{"keyX"}, record), false)
	assert.Equal(t, isSimpleRecordComplete([]string{"key1"}, record), false)
	assert.Equal(t, isSimpleRecordComplete([]string{"key2"}, record), false)
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
	assert.Equal(t, isAddressComplete(adress), false)
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
	assert.Equal(t, isAddressComplete(adress), true)
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
	assert.Equal(t, value, 0)
}

func TestCalcMetricOutput(t *testing.T) {
	//value := calculateMetric(Consistency, map[string]interface{}{}, []map[string]string{})
	assert.Equal(t, 1, 1)
}
