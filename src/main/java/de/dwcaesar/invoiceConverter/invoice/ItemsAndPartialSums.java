package de.dwcaesar.invoiceConverter.invoice;

import de.dwcaesar.invoiceConverter.invoice.model.Item;

import java.math.BigDecimal;
import java.util.List;

public class ItemsAndPartialSums {
    private final List<Item> items;
    private final BigDecimal partialSumFull;
    private final BigDecimal partialSumReduced;

    public ItemsAndPartialSums(List<Item> items, BigDecimal partialSumFull, BigDecimal partialSumReduced) {
        this.items = items;
        this.partialSumFull = partialSumFull;
        this.partialSumReduced = partialSumReduced;
    }

    public List<Item> getItems() {
        return items;
    }

    public BigDecimal getPartialSumFull() {
        return partialSumFull;
    }

    public BigDecimal getPartialSumReduced() {
        return partialSumReduced;
    }

}