package de.dwcaesar.invoiceConverter.invoice;

import de.dwcaesar.invoiceConverter.invoice.model.VATType;
import de.dwcaesar.invoiceConverter.io.model.Vat;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.math.RoundingMode;
import java.util.Map;

/**
 * Implementation of the VatConverter interface
 */
@Service
public class VatConverterImpl implements VatConverter {

    @Override
    public Map<VATType, BigDecimal> mapVatRates(Map<Vat, BigDecimal> vat) {
        return Map.of(VATType.FULL, vat.getOrDefault(Vat.FULL, BigDecimal.ZERO),
                VATType.REDUCED, vat.getOrDefault(Vat.REDUCED, BigDecimal.ZERO));
    }

    @Override
    public BigDecimal calculateVatPortion(BigDecimal itemPrice, BigInteger amount, BigDecimal vat) {
        // Mehrwertsteueranteil: Netto * VAT = (Brutto * VAT / (1 + VAT))
        return itemPrice.multiply(BigDecimal.valueOf(amount.doubleValue())).multiply(vat)
                .divide(BigDecimal.ONE.add(vat), 2, RoundingMode.HALF_UP);
    }
}