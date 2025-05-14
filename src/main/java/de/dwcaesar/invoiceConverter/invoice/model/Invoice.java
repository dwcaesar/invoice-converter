package de.dwcaesar.invoiceConverter.invoice.model;

import io.swagger.v3.oas.annotations.media.Schema;
import lombok.Builder;
import lombok.Data;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;

@Data
@Builder
@Schema(description = "Invoice data in JSON format (converted from XML)")
public class Invoice {
    @Schema(description = "Unique invoice number", example = "12345")
    private Integer invoiceNumber;

    @Schema(description = "Billing address information")
    private Address billingAddress;

    @Schema(description = "Shipping address information")
    private Address shippingAddress;

    @Schema(description = "Method of payment for this invoice")
    private PaymentMethod paymentMethod;

    @Schema(description = "List of items included in this invoice")
    private List<Item> items;

    @Schema(description = "Partial sum for items with full VAT rate", example = "80.00")
    private BigDecimal partialSumFull;

    @Schema(description = "Partial sum for items with reduced VAT rate", example = "20.00")
    private BigDecimal partialSumReduced;

    @Schema(description = "Net total amount of the invoice", example = "100.00")
    private BigDecimal totalSumNetto;

    @Schema(description = "Gross total amount of the invoice", example = "119.00")
    private BigDecimal totalSumBrutto;

    @Schema(description = "VAT amounts by VAT type")
    private Map<VATType, BigDecimal> vatRates;
}
