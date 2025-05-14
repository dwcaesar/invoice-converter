package de.dwcaesar.invoiceConverter.invoice;

import de.dwcaesar.invoiceConverter.invoice.model.VATType;
import de.dwcaesar.invoiceConverter.io.model.Vat;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.util.Map;

/**
 * Service for handling VAT rate mapping and calculations
 */
public interface VatConverter {
    
    /**
     * Maps IO VAT rates to domain VAT rates
     * @param vat the IO VAT rates
     * @return the mapped domain VAT rates
     */
    Map<VATType, BigDecimal> mapVatRates(Map<Vat, BigDecimal> vat);
    
    /**
     * Calculates the VAT portion of an item
     * @param itemPrice the price of the item
     * @param amount the amount of the item
     * @param vat the VAT rate
     * @return the VAT portion
     */
    BigDecimal calculateVatPortion(BigDecimal itemPrice, BigInteger amount, BigDecimal vat);
}