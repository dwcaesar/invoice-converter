package de.dwcaesar.invoiceConverter.invoice;

import de.dwcaesar.invoiceConverter.invoice.model.*;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

/**
 * Service for converting between IO and domain invoice models
 */
@Service
@RequiredArgsConstructor
public class InvoiceConverter {

    private final AddressConverter addressConverter;
    private final VatConverter vatConverter;
    private final ItemConverter itemConverter;

    /**
     * Converts an IO invoice to a domain invoice
     * @param invoice the IO invoice to convert
     * @return the converted domain invoice
     */
    public Invoice convert(de.dwcaesar.invoiceConverter.io.model.Invoice invoice) {
        Address shippingAddress = addressConverter.convertShippingAddress(invoice.getShippingAddress());
        Address billingAddress = addressConverter.convertBillingAddress(invoice.getBillingAddress())
                .orElse(shippingAddress);
        ItemsAndPartialSums itemsAndPartialSums = itemConverter.convertItemsAndCalculatePartialSums(
                invoice.getItems(), invoice.getVat());

        return Invoice.builder()
                .invoiceNumber(invoice.getInvoiceNumber().intValue())
                .billingAddress(billingAddress)
                .shippingAddress(shippingAddress)
                .paymentMethod(switch (invoice.getPaymentMethod()) {
                    case CREDIT_CARD -> PaymentMethod.CREDIT_CARD;
                    case INVOICE -> PaymentMethod.INVOICE;
                    case SURNAME -> PaymentMethod.SURNAME;
                })
                .items(itemsAndPartialSums.getItems())
                .totalSumNetto(invoice.getNetto())
                .totalSumBrutto(invoice.getBrutto())
                .partialSumFull(itemsAndPartialSums.getPartialSumFull())
                .partialSumReduced(itemsAndPartialSums.getPartialSumReduced())
                .vatRates(vatConverter.mapVatRates(invoice.getVat()))
                .build();
    }
}
