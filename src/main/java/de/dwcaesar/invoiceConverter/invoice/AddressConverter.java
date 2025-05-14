package de.dwcaesar.invoiceConverter.invoice;

import de.dwcaesar.invoiceConverter.invoice.model.Address;

import java.util.Optional;

/**
 * Service for converting between IO and domain address models
 */
public interface AddressConverter {
    
    /**
     * Converts an IO address to a domain address
     * @param shippingAddress the IO address to convert
     * @return the converted domain address
     */
    Address convertShippingAddress(de.dwcaesar.invoiceConverter.io.model.Address shippingAddress);
    
    /**
     * Converts an IO billing address to an optional domain address
     * @param billingAddress the IO billing address to convert
     * @return an Optional containing the converted domain address, or empty if the input was null
     */
    Optional<Address> convertBillingAddress(de.dwcaesar.invoiceConverter.io.model.Address billingAddress);
}