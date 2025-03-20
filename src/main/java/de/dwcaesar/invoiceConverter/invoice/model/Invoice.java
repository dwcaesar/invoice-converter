package de.dwcaesar.invoiceConverter.invoice.model;

import lombok.Builder;
import lombok.Data;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;

@Data
@Builder
public class Invoice {
    private Integer invoiceNumber;
    private Address billingAddress;
    private Address shippingAddress;
    private PaymentMethod paymentMethod;
    private List<Item> items;
    private BigDecimal partialSumFull;
    private BigDecimal partialSumReduced;
    private BigDecimal totalSumNetto;
    private BigDecimal totalSumBrutto;
    private Map<VATType, BigDecimal> vatRates;
}
