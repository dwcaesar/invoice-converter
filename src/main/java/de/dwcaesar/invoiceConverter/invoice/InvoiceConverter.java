package de.dwcaesar.invoiceConverter.invoice;

import org.springframework.stereotype.Service;

import java.math.BigDecimal;

@Service
public class InvoiceConverter {

    public Invoice convert(generated.Invoice invoice) {
        return Invoice.builder()
                .invoiceNumber(invoice.getInvoiceNumber().intValue())
                .billingAddress(Address.builder()
                        .name(invoice.getBillingAddress().getName()) // FIXME - breaks when billing address is missing
                        .street(invoice.getBillingAddress().getStreet())
                        .city(invoice.getBillingAddress().getCity())
                        .zipcode(invoice.getBillingAddress().getZip())
                        .build())
                .shippingAddress(Address.builder()
                        .name(invoice.getShippingAddress().getName())
                        .street(invoice.getShippingAddress().getStreet())
                        .city(invoice.getShippingAddress().getCity())
                        .zipcode(invoice.getShippingAddress().getZip())
                        .build())
                .paymentMethod(switch (invoice.getPaymentMethod()) {
                    case CREDIT_CARD -> PaymentMethod.CREDIT_CARD;
                    case INVOICE -> PaymentMethod.INVOICE;
                    case SURNAME -> PaymentMethod.SURNAME;
                })
                .items(invoice.getItem().stream()
                        .map(item -> Item.builder()
                                .name(item.getName())
                                .amount(item.getAmount().intValue())
                                .itemPrice(item.getItemPrice())
                                .vatType(switch (item.getVat()) {
                                    case FULL -> VATType.FULL;
                                    case REDUCED -> VATType.REDUCED;
                                    case NONE -> VATType.NONE;
                                })
                                .positionSum(item.getItemPrice().multiply(BigDecimal.valueOf(item.getAmount().doubleValue())))
                                .build())
                        .toList())
                .totalSumNetto(invoice.getNetto())
                .totalSumBrutto(invoice.getBrutto())
                .partialSumFull(BigDecimal.ZERO) // TODO - get from items and local constant
                .partialSumReduced(BigDecimal.ZERO) // TODO - get from items and local constant
                .build(); // TODO - validate: all required fields set? do sums match?
    }
}
