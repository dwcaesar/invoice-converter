package de.dwcaesar.invoiceConverter.invoice;

import de.dwcaesar.invoiceConverter.io.model.Vat;

import java.util.List;
import java.util.Map;

/**
 * Service for converting items and calculating partial sums
 */
public interface ItemConverter {
    
    /**
     * Converts IO items to domain items and calculates partial sums
     * @param items the IO items to convert
     * @param vat the VAT rates
     * @return an ItemsAndPartialSums object containing the converted items and partial sums
     */
    ItemsAndPartialSums convertItemsAndCalculatePartialSums(
            List<de.dwcaesar.invoiceConverter.io.model.Item> items, 
            Map<Vat, java.math.BigDecimal> vat);
}