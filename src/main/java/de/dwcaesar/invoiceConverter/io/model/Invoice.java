package de.dwcaesar.invoiceConverter.io.model;

import lombok.Getter;
import lombok.Setter;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.util.List;
import java.util.Map;

@Getter
@Setter
public class Invoice {
private BigInteger invoiceNumber;
    private Address billingAddress;
    private Address shippingAddress;
    private PaymentMethod paymentMethod;
    private List<Item> items;
    private BigDecimal netto;
    private BigDecimal brutto;
    private Map<Vat, BigDecimal> vat;
}
