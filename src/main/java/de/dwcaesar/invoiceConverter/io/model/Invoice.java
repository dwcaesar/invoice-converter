package de.dwcaesar.invoiceConverter.io.model;

import io.swagger.v3.oas.annotations.media.Schema;
import lombok.Getter;
import lombok.Setter;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.util.List;
import java.util.Map;

@Getter
@Setter
@Schema(description = "Invoice data in XML format")
public class Invoice {
    @Schema(description = "Unique invoice number", example = "12345")
    private BigInteger invoiceNumber;

    @Schema(description = "Billing address information")
    private Address billingAddress;

    @Schema(description = "Shipping address information")
    private Address shippingAddress;

    @Schema(description = "Method of payment for this invoice")
    private PaymentMethod paymentMethod;

    @Schema(description = "List of items included in this invoice")
    private List<Item> items;

    @Schema(description = "Net total amount of the invoice", example = "100.00")
    private BigDecimal netto;

    @Schema(description = "Gross total amount of the invoice", example = "119.00")
    private BigDecimal brutto;

    @Schema(description = "VAT amounts by VAT type")
    private Map<Vat, BigDecimal> vat;
}
