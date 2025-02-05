package de.dwcaesar.invoiceConverter.io.model;

import lombok.Getter;
import lombok.Setter;

import java.math.BigDecimal;
import java.math.BigInteger;

@Getter
@Setter
public class Item {
    private String name;
    private BigInteger amount;
    private BigDecimal itemPrice;
    private Vat vat;
}
