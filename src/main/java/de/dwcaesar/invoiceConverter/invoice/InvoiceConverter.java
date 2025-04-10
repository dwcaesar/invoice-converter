package de.dwcaesar.invoiceConverter.invoice;

import de.dwcaesar.invoiceConverter.invoice.model.*;
import de.dwcaesar.invoiceConverter.io.model.Vat;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.math.RoundingMode;
import java.util.*;

@Service
public class InvoiceConverter {

    public Invoice convert(de.dwcaesar.invoiceConverter.io.model.Invoice invoice) {
        Address shippingAddress = getShippingAddress(invoice.getShippingAddress());
        Address billingAddress = getBillingAddress(invoice.getBillingAddress())
                .orElse(shippingAddress);
        ItemsAndPartialSums itemsAndPartialSums = getItemsAndPartialSums(invoice.getItems(), invoice.getVat());

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
                .vatRates(mapVatRates(invoice.getVat()))
                .build();
    }

    private Map<VATType, BigDecimal> mapVatRates(Map<Vat, BigDecimal> vat) {
        return Map.of(VATType.FULL, vat.getOrDefault(Vat.FULL, BigDecimal.ZERO),
                VATType.REDUCED, vat.getOrDefault(Vat.REDUCED, BigDecimal.ZERO));
    }

    private Address getShippingAddress(de.dwcaesar.invoiceConverter.io.model.Address shippingAddress) {
        return Address.builder()
                .name(shippingAddress.getName())
                .street(shippingAddress.getStreet())
                .city(shippingAddress.getCity())
                .zipcode(shippingAddress.getZip())
                .build();
    }

    private Optional<Address> getBillingAddress(de.dwcaesar.invoiceConverter.io.model.Address billingAddress) {
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

    private ItemsAndPartialSums getItemsAndPartialSums(List<de.dwcaesar.invoiceConverter.io.model.Item> items, Map<Vat, BigDecimal> vat) {
        BigDecimal partialSumFull = BigDecimal.ZERO;
        BigDecimal partialSumReduced = BigDecimal.ZERO;
        List<Item> convertedItems = new ArrayList<>();
        for (de.dwcaesar.invoiceConverter.io.model.Item item : items) {
            VATType vatType = switch (item.getVat()) {
                case FULL -> {
                    partialSumFull = partialSumFull.add(toVatPortion(item.getItemPrice(), item.getAmount(), vat.get(Vat.FULL)));
                    yield VATType.FULL;
                }
                case REDUCED -> {
                    partialSumReduced = partialSumReduced.add(toVatPortion(item.getItemPrice(), item.getAmount(), vat.get(Vat.REDUCED)));
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
