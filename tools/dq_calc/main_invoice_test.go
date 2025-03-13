package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInvoicesExist(t *testing.T) {
	_, err := parseInvoiceXML("assets/invoices.xml")

	//Fails immediately if error is returned
	require.NoError(t, err, "Failed to parse invoices.xml")
}

func TestLoadInvoiceXML(t *testing.T) {
	invoices, err := parseInvoiceXML("assets/invoices.xml")
	require.NoError(t, err, "Failed to parse invoices.xml")

	require.NotNil(t, invoices, "Invoices should not be nil")
	assert.Equal(t, 2, len(invoices), "invoices length not matching %s", invoices)

	for i, invoice := range invoices {
		assert.NotNil(t, invoice, "Invoice at index %d should not be nil", i)
		require.NotEmpty(t, invoice.InvoiceNumber, "Invoice at index %d should have an invoice number", i)
	}
}

func TestConvertInvoiceToMap(t *testing.T) {
	// Sample full invoice object
	invoice := Invoice{
		InvoiceNumber: "12345",
		BillingAddress: Address{
			Name:   "John Doe",
			Street: "Main St 1",
			Zip:    "12345",
			City:   "Sample City",
		},
		ShippingAddress: Address{
			Name:   "Jane Doe",
			Street: "Second St 2",
			Zip:    "67890",
			City:   "Another City",
		},
		PaymentMethod: "Credit Card",
		Items: []Item{
			{Name: "Product A", Amount: "2", ItemPrice: "50.00", Vat: "full"},
			{Name: "Product B", Amount: "1", ItemPrice: "30.00", Vat: "reduced"},
		},
		Netto:  "130.00",
		Brutto: "154.70",
	}

	invoiceMap, err := convertInvoiceToMap(invoice)
	if err != nil {
		t.Fatalf("Error converting invoice to map, %v", err)
	}

	expectedKeys := []string{"InvoiceNumber", "BillingAddress", "ShippingAddress", "PaymentMethod", "Items", "Netto", "Brutto"}

	for _, key := range expectedKeys {
		assert.NotNil(t, invoiceMap[key], "%v does not exist in the map", key)
	}

	// Check nested fields
	billingAddress, ok := invoiceMap["BillingAddress"].(map[string]interface{})
	require.Truef(t, ok, "Failed to convertbillingAddress to map")

	expectedBillingAddressKeys := []string{"Name", "Street", "Zip", "City"}
	for _, key := range expectedBillingAddressKeys {
		assert.NotNil(t, billingAddress[key], "%v does not exist in the billingAddress map", key)
	}

	// Check list items
	items, ok := invoiceMap["Items"].([]interface{})
	require.Truef(t, ok, "Failed to convert items to a map")
	assert.Equal(t, 2, len(items), "Expected 2 items, but got %d", len(items))

	// Check item structure
	firstItem, ok := items[0].(map[string]interface{})
	require.Truef(t, ok, "Failed to convert an item to a map")

	expectedItemKeys := []string{"Name", "Amount", "ItemPrice", "Vat"}
	for _, key := range expectedItemKeys {
		assert.NotNil(t, firstItem[key], "%v does not exist in the item map", key)
	}
}

func TestFalseInvoiceCompleteness(t *testing.T) {
	invoice := map[string]interface{}{
		"invoiceNumber": "123456",
		"billingAddress": map[string]interface{}{
			"name":   "John Doe",
			"street": "Main Street 1",
			"zip":    "12345",
			"city":   "Sample City",
		},
	}
	var fields = map[string]interface{}{
		"invoiceNumber":  true,
		"billingAddress": []interface{}{"name", "street", "zip", "city"},
		"netto":          true,
	}

	result := validateInvoiceCompleteness(invoice, fields)

	assert.False(t, result, "Expected invoice to be invalid")
}

func TestSimpleInvoiceCompleteness(t *testing.T) {
	invoice := map[string]interface{}{
		"invoiceNumber": "123456",
		"billingAddress": map[string]interface{}{
			"name":   "John Doe",
			"street": "Main Street 1",
			"zip":    "12345",
			"city":   "Sample City",
		},
		"netto": "1000",
	}
	var fields = map[string]interface{}{
		"invoiceNumber":  true,
		"billingAddress": []interface{}{"name", "street", "zip", "city"},
		"netto":          true,
	}

	result := validateInvoiceCompleteness(invoice, fields)

	assert.True(t, result, "Expected invoice to be valid")
}
func TestFalseComplexInvoiceCompleteness(t *testing.T) {
	invoice := map[string]interface{}{
		"invoiceNumber": "123456",
		"billingAddress": map[string]interface{}{
			"name":   "John Doe",
			"street": "Main Street 1",
			"zip":    "12345",
		},
		"netto": "1000",
	}
	var fields = map[string]interface{}{
		"invoiceNumber":  true,
		"billingAddress": []interface{}{"name", "street", "zip", "city"},
		"netto":          true,
	}

	result := validateInvoiceCompleteness(invoice, fields)

	assert.False(t, result, "Expected invoice to be invalid")
}
func TestComplexInvoiceCompleteness(t *testing.T) {
	invoice := map[string]interface{}{
		"invoiceNumber": "123456",
		"billingAddress": map[string]interface{}{
			"name":   "John Doe",
			"street": "Main Street 1",
			"zip":    "12345",
			"city":   "Main City",
		},
		"shippingAddress": map[string]interface{}{
			"name":   "John Doe",
			"street": "Second Street 2",
			"zip":    "67890",
			"city":   "Another City",
		},
		"paymentMethod": "credit-card",
		"items": []interface{}{
			map[string]interface{}{
				"name":      "Product A",
				"amount":    "1",
				"itemPrice": "50.00",
				"vat":       "full",
			},
			map[string]interface{}{
				"name":      "Product B",
				"amount":    "2",
				"itemPrice": "80.00",
				"vat":       "full",
			},
		},
		"netto":  "1000",
		"brutto": "154.70",
	}
	var fields = map[string]interface{}{
		"invoiceNumber": true,
		"billingAddress": []interface{}{
			"name", "street", "zip", "city",
		},
		"shippingAddress": []interface{}{
			"name", "street", "zip", "city",
		},
		"paymentMethod": true,
		"items": []interface{}{
			"name", "amount", "itemPrice", "vat",
		},
		"netto":  true,
		"brutto": true,
	}

	result := validateInvoiceCompleteness(invoice, fields)

	assert.True(t, result, "Expected invoice to be valid")
}

func TestFalseItemsInvoiceCompleteness(t *testing.T) {
	invoice := map[string]interface{}{
		"invoiceNumber": "123456",
		"billingAddress": map[string]interface{}{
			"name":   "John Doe",
			"street": "Main Street 1",
			"zip":    "12345",
			"city":   "Main City",
		},
		"items": []interface{}{
			map[string]interface{}{
				"name":      "Product A",
				"amount":    "1",
				"itemPrice": "50.00",
				"vat":       "full",
			},
			map[string]interface{}{
				"name":      "Product B",
				"amount":    "2",
				"itemPrice": "80.00",
			},
		},
	}
	var fields = map[string]interface{}{
		"invoiceNumber": true,
		"billingAddress": []interface{}{
			"name", "street", "zip", "city",
		},
		"items": []interface{}{
			"name", "amount", "itemPrice", "vat",
		},
	}

	result := validateInvoiceCompleteness(invoice, fields)

	assert.False(t, result, "Expected invoice to be invalid")
}
