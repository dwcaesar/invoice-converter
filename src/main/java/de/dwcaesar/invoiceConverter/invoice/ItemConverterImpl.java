package de.dwcaesar.invoiceConverter.invoice;

import de.dwcaesar.invoiceConverter.invoice.model.Item;
import de.dwcaesar.invoiceConverter.invoice.model.VATType;
import de.dwcaesar.invoiceConverter.io.model.Vat;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

/**
 * Implementation of the ItemConverter interface
 */
@Service
@RequiredArgsConstructor
public class ItemConverterImpl implements ItemConverter {

    private final VatConverter vatConverter;

    @Override
    public ItemsAndPartialSums convertItemsAndCalculatePartialSums(
            List<de.dwcaesar.invoiceConverter.io.model.Item> items,
            Map<Vat, BigDecimal> vat) {

        BigDecimal partialSumFull = BigDecimal.ZERO;
        BigDecimal partialSumReduced = BigDecimal.ZERO;
        List<Item> convertedItems = new ArrayList<>();

        for (de.dwcaesar.invoiceConverter.io.model.Item item : items) {
            VATType vatType = switch (item.getVat()) {
                case FULL -> {
                    partialSumFull = partialSumFull.add(
                            vatConverter.calculateVatPortion(item.getItemPrice(), item.getAmount(), vat.get(Vat.FULL)));
                    yield VATType.FULL;
                }
                case REDUCED -> {
                    partialSumReduced = partialSumReduced.add(
                            vatConverter.calculateVatPortion(item.getItemPrice(), item.getAmount(), vat.get(Vat.REDUCED)));
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
}