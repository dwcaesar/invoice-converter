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
