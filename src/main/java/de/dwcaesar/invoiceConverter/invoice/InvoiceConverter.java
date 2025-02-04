package de.dwcaesar.invoiceConverter.invoice;

import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.math.RoundingMode;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@Service
public class InvoiceConverter {

    private final BigDecimal VAT_FULL = BigDecimal.valueOf(19, 2);
    private final BigDecimal VAT_REDUCED = BigDecimal.valueOf(7, 2);

    public Invoice convert(generated.Invoice invoice) {
        Address shippingAddress = getShippingAddress(invoice.getShippingAddress());
        Address billingAddress = getBillingAddress(invoice.getBillingAddress())
                .orElse(shippingAddress);
        ItemsAndPartialSums itemsAndPartialSums = getItemsAndPartialSums(invoice.getItem());

        return Invoice.builder()
                .invoiceNumber(invoice.getInvoiceNumber().intValue())
                .billingAddress(billingAddress)
                .shippingAddress(shippingAddress)
                .paymentMethod(switch (invoice.getPaymentMethod()) {
                    case CREDIT_CARD -> PaymentMethod.CREDIT_CARD;
                    case INVOICE -> PaymentMethod.INVOICE;
                    case SURNAME -> PaymentMethod.SURNAME;
                })
                .items(itemsAndPartialSums.items)
                .totalSumNetto(invoice.getNetto())
                .totalSumBrutto(invoice.getBrutto())
                .partialSumFull(itemsAndPartialSums.partialSumFull)
                .partialSumReduced(itemsAndPartialSums.partialSumReduced)
                .build();
    }

    private Address getShippingAddress(generated.Address shippingAddress) {
        return Address.builder()
                .name(shippingAddress.getName())
                .street(shippingAddress.getStreet())
                .city(shippingAddress.getCity())
                .zipcode(shippingAddress.getZip())
                .build();
    }

    private Optional<Address> getBillingAddress(generated.Address billingAddress) {
        if (null == billingAddress) {
            return Optional.empty();
        } else {
            return Optional.of(Address.builder()
                    .name(billingAddress.getName())
                    .street(billingAddress.getStreet())
                    .city(billingAddress.getCity())
                    .zipcode(billingAddress.getZip())
                    .build());
        }
    }

    private ItemsAndPartialSums getItemsAndPartialSums(List<generated.Item> items) {
        BigDecimal partialSumFull = BigDecimal.ZERO;
        BigDecimal partialSumReduced = BigDecimal.ZERO;
        List<Item> convertedItems = new ArrayList<>();
        for (generated.Item item : items) {
            VATType vatType = switch (item.getVat()) {
                case FULL -> {
                    partialSumFull = partialSumFull.add(toVatPortion(item.getItemPrice(), item.getAmount(), VAT_FULL));
                    yield VATType.FULL;
                }
                case REDUCED -> {
                    partialSumReduced = partialSumReduced.add(toVatPortion(item.getItemPrice(), item.getAmount(), VAT_REDUCED));
                    yield VATType.REDUCED;
                }
                case NONE -> VATType.NONE;
            };
            convertedItems.add(Item.builder()
                    .name(item.getName())
                    .amount(item.getAmount().intValue())
                    .itemPrice(item.getItemPrice())
                    .vatType(vatType)
                    .positionSum(item.getItemPrice().multiply(BigDecimal.valueOf(item.getAmount().doubleValue())))
                    .build());
        }

        return new ItemsAndPartialSums(convertedItems, partialSumFull, partialSumReduced);
    }

    private BigDecimal toVatPortion(BigDecimal itemPrice, BigInteger amount, BigDecimal vat) {
        // Mehrwertsteueranteil: Netto * VAT = (Brutto * VAT / (1 + VAT))
        return itemPrice.multiply(BigDecimal.valueOf(amount.doubleValue())).multiply(vat).divide(BigDecimal.ONE.add(vat), 2, RoundingMode.HALF_UP);
    }

    private record ItemsAndPartialSums(List<Item> items, BigDecimal partialSumFull, BigDecimal partialSumReduced) {
    }
}
