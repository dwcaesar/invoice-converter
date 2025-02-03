package de.dwcaesar.invoiceConverter.invoice;

import lombok.Builder;
import lombok.Data;

import java.math.BigDecimal;

@Data
@Builder
public class Item {
    private String name;
    private Integer amount;
    private BigDecimal itemPrice;
    private BigDecimal positionSum;
}
