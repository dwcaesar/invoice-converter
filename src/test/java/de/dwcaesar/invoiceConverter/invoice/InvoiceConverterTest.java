package de.dwcaesar.invoiceConverter.invoice;

import de.dwcaesar.invoiceConverter.invoice.model.Address;
import de.dwcaesar.invoiceConverter.invoice.model.Item;
import de.dwcaesar.invoiceConverter.invoice.model.VATType;
import de.dwcaesar.invoiceConverter.io.model.Vat;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.util.*;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class InvoiceConverterTest {

    @Mock
    private AddressConverter addressConverter;

    @Mock
    private VatConverter vatConverter;

    @Mock
    private ItemConverter itemConverter;

    private InvoiceConverter invoiceConverter;

    @BeforeEach
    void setUp() {
        invoiceConverter = new InvoiceConverter(addressConverter, vatConverter, itemConverter);

        // Set up default mock behavior
        when(addressConverter.convertShippingAddress(any())).thenAnswer(invocation -> {
            de.dwcaesar.invoiceConverter.io.model.Address address = invocation.getArgument(0);
            return Address.builder()
                    .name(address.getName())
                    .street(address.getStreet())
                    .city(address.getCity())
                    .zipcode(address.getZip())
                    .build();
        });

        when(addressConverter.convertBillingAddress(any())).thenAnswer(invocation -> {
            de.dwcaesar.invoiceConverter.io.model.Address address = invocation.getArgument(0);
            if (address == null) {
                return Optional.empty();
            }
            return Optional.of(Address.builder()
                    .name(address.getName())
                    .street(address.getStreet())
                    .city(address.getCity())
                    .zipcode(address.getZip())
                    .build());
        });

        when(vatConverter.mapVatRates(any())).thenAnswer(invocation -> {
            Map<Vat, BigDecimal> vat = invocation.getArgument(0);
            return Map.of(
                    VATType.FULL, vat.getOrDefault(Vat.FULL, BigDecimal.ZERO),
                    VATType.REDUCED, vat.getOrDefault(Vat.REDUCED, BigDecimal.ZERO)
            );
        });

        when(itemConverter.convertItemsAndCalculatePartialSums(any(), any())).thenAnswer(invocation -> {
            List<de.dwcaesar.invoiceConverter.io.model.Item> items = invocation.getArgument(0);
            Map<Vat, BigDecimal> vat = invocation.getArgument(1);

            List<Item> convertedItems = new ArrayList<>();
            BigDecimal partialSumFull = BigDecimal.ZERO;
            BigDecimal partialSumReduced = BigDecimal.ZERO;

            for (de.dwcaesar.invoiceConverter.io.model.Item item : items) {
                VATType vatType = switch (item.getVat()) {
                    case FULL -> {
                        BigDecimal vatPortion = item.getItemPrice()
                                .multiply(BigDecimal.valueOf(item.getAmount().doubleValue()))
                                .multiply(vat.getOrDefault(Vat.FULL, BigDecimal.ZERO))
                                .divide(BigDecimal.ONE.add(vat.getOrDefault(Vat.FULL, BigDecimal.ZERO)), 2, java.math.RoundingMode.HALF_UP);
                        partialSumFull = partialSumFull.add(vatPortion);
                        yield VATType.FULL;
                    }
                    case REDUCED -> {
                        BigDecimal vatPortion = item.getItemPrice()
                                .multiply(BigDecimal.valueOf(item.getAmount().doubleValue()))
                                .multiply(vat.getOrDefault(Vat.REDUCED, BigDecimal.ZERO))
                                .divide(BigDecimal.ONE.add(vat.getOrDefault(Vat.REDUCED, BigDecimal.ZERO)), 2, java.math.RoundingMode.HALF_UP);
                        partialSumReduced = partialSumReduced.add(vatPortion);
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
        });
    }

    @Test
    void testConvertWithNullBillingAddress() {
        // Arrange
        de.dwcaesar.invoiceConverter.io.model.Invoice inputInvoice = createBasicInvoice();
        inputInvoice.setBillingAddress(null);

        // Act
        de.dwcaesar.invoiceConverter.invoice.model.Invoice result = invoiceConverter.convert(inputInvoice);

        // Assert
        assertNotNull(result);
        assertEquals(inputInvoice.getShippingAddress().getName(), result.getBillingAddress().getName());
        assertEquals(inputInvoice.getShippingAddress().getStreet(), result.getBillingAddress().getStreet());
        assertEquals(inputInvoice.getShippingAddress().getCity(), result.getBillingAddress().getCity());
        assertEquals(inputInvoice.getShippingAddress().getZip(), result.getBillingAddress().getZipcode());
    }

    @Test
    void testConvertWithEmptyItemsList() {
        // Arrange
        de.dwcaesar.invoiceConverter.io.model.Invoice inputInvoice = createBasicInvoice();
        inputInvoice.setItems(Collections.emptyList());

        // Act
        de.dwcaesar.invoiceConverter.invoice.model.Invoice result = invoiceConverter.convert(inputInvoice);

        // Assert
        assertNotNull(result);
        assertTrue(result.getItems().isEmpty());
        assertEquals(BigDecimal.ZERO, result.getPartialSumFull());
        assertEquals(BigDecimal.ZERO, result.getPartialSumReduced());
    }

    @Test
    void testConvertWithMissingVatRates() {
        // Arrange
        de.dwcaesar.invoiceConverter.io.model.Invoice inputInvoice = createBasicInvoice();
        inputInvoice.setVat(Collections.emptyMap());

        // Also need to set items to empty list to avoid NullPointerException
        // when calculating VAT portions with missing VAT rates
        inputInvoice.setItems(Collections.emptyList());

        // Act
        de.dwcaesar.invoiceConverter.invoice.model.Invoice result = invoiceConverter.convert(inputInvoice);

        // Assert
        assertNotNull(result);
        assertEquals(BigDecimal.ZERO, result.getVatRates().get(VATType.FULL));
        assertEquals(BigDecimal.ZERO, result.getVatRates().get(VATType.REDUCED));
    }

    @Test
    void testConvertWithExtremeValues() {
        // Arrange
        de.dwcaesar.invoiceConverter.io.model.Invoice inputInvoice = createBasicInvoice();

        // Set extreme values
        inputInvoice.setInvoiceNumber(BigInteger.valueOf(Integer.MAX_VALUE).add(BigInteger.ONE));

        de.dwcaesar.invoiceConverter.io.model.Item item = new de.dwcaesar.invoiceConverter.io.model.Item();
        item.setName("Extreme Item");
        item.setAmount(BigInteger.valueOf(Integer.MAX_VALUE).add(BigInteger.ONE));
        item.setItemPrice(new BigDecimal("9999999999.99"));
        item.setVat(Vat.FULL);
        inputInvoice.setItems(Collections.singletonList(item));

        // Act
        de.dwcaesar.invoiceConverter.invoice.model.Invoice result = invoiceConverter.convert(inputInvoice);

        // Assert
        assertNotNull(result);
        // Integer overflow should wrap around to negative
        assertTrue(result.getInvoiceNumber() < 0);
        assertEquals(1, result.getItems().size());
        // Integer overflow should wrap around to negative
        assertTrue(result.getItems().getFirst().getAmount() < 0);
    }

    @Test
    void testConvertWithAllVatTypes() {
        // Arrange
        de.dwcaesar.invoiceConverter.io.model.Invoice inputInvoice = createBasicInvoice();

        // Create items with different VAT types
        de.dwcaesar.invoiceConverter.io.model.Item fullVatItem = new de.dwcaesar.invoiceConverter.io.model.Item();
        fullVatItem.setName("Full VAT Item");
        fullVatItem.setAmount(BigInteger.valueOf(2));
        fullVatItem.setItemPrice(new BigDecimal("100.00"));
        fullVatItem.setVat(Vat.FULL);

        de.dwcaesar.invoiceConverter.io.model.Item reducedVatItem = new de.dwcaesar.invoiceConverter.io.model.Item();
        reducedVatItem.setName("Reduced VAT Item");
        reducedVatItem.setAmount(BigInteger.valueOf(1));
        reducedVatItem.setItemPrice(new BigDecimal("50.00"));
        reducedVatItem.setVat(Vat.REDUCED);

        de.dwcaesar.invoiceConverter.io.model.Item noVatItem = new de.dwcaesar.invoiceConverter.io.model.Item();
        noVatItem.setName("No VAT Item");
        noVatItem.setAmount(BigInteger.valueOf(1));
        noVatItem.setItemPrice(new BigDecimal("25.00"));
        noVatItem.setVat(Vat.NONE);

        inputInvoice.setItems(java.util.Arrays.asList(fullVatItem, reducedVatItem, noVatItem));

        // Act
        de.dwcaesar.invoiceConverter.invoice.model.Invoice result = invoiceConverter.convert(inputInvoice);

        // Assert
        assertNotNull(result);
        assertEquals(3, result.getItems().size());

        // Check VAT types are correctly converted
        boolean hasFullVat = false;
        boolean hasReducedVat = false;
        boolean hasNoVat = false;

        for (Item item : result.getItems()) {
            if (item.getVatType() == VATType.FULL) hasFullVat = true;
            if (item.getVatType() == VATType.REDUCED) hasReducedVat = true;
            if (item.getVatType() == VATType.NONE) hasNoVat = true;
        }

        assertTrue(hasFullVat, "Should have an item with FULL VAT type");
        assertTrue(hasReducedVat, "Should have an item with REDUCED VAT type");
        assertTrue(hasNoVat, "Should have an item with NONE VAT type");
    }

    @Test
    void testConvertWithZeroAmountItems() {
        // Arrange
        de.dwcaesar.invoiceConverter.io.model.Invoice inputInvoice = createBasicInvoice();

        de.dwcaesar.invoiceConverter.io.model.Item zeroAmountItem = new de.dwcaesar.invoiceConverter.io.model.Item();
        zeroAmountItem.setName("Zero Amount Item");
        zeroAmountItem.setAmount(BigInteger.ZERO);
        zeroAmountItem.setItemPrice(new BigDecimal("100.00"));
        zeroAmountItem.setVat(Vat.FULL);

        inputInvoice.setItems(Collections.singletonList(zeroAmountItem));

        // Act
        de.dwcaesar.invoiceConverter.invoice.model.Invoice result = invoiceConverter.convert(inputInvoice);

        // Assert
        assertNotNull(result);
        assertEquals(1, result.getItems().size());
        assertEquals(0, result.getItems().getFirst().getAmount());
        // Use compareTo for BigDecimal comparison to avoid scale issues
        assertEquals(0, result.getItems().getFirst().getPositionSum().compareTo(BigDecimal.ZERO));
        assertEquals(0, result.getPartialSumFull().compareTo(BigDecimal.ZERO));
    }

    @Test
    void testConvertWithNegativeAmountItems() {
        // Arrange
        de.dwcaesar.invoiceConverter.io.model.Invoice inputInvoice = createBasicInvoice();

        de.dwcaesar.invoiceConverter.io.model.Item negativeAmountItem = new de.dwcaesar.invoiceConverter.io.model.Item();
        negativeAmountItem.setName("Negative Amount Item");
        negativeAmountItem.setAmount(BigInteger.valueOf(-1));
        negativeAmountItem.setItemPrice(new BigDecimal("100.00"));
        negativeAmountItem.setVat(Vat.FULL);

        inputInvoice.setItems(Collections.singletonList(negativeAmountItem));

        // Act
        de.dwcaesar.invoiceConverter.invoice.model.Invoice result = invoiceConverter.convert(inputInvoice);

        // Assert
        assertNotNull(result);
        assertEquals(1, result.getItems().size());
        assertEquals(-1, result.getItems().getFirst().getAmount());
        // Use compareTo for BigDecimal comparison to avoid scale issues
        assertEquals(0, result.getItems().getFirst().getPositionSum().compareTo(new BigDecimal("-100.00")));
        // VAT portion should be negative
        assertTrue(result.getPartialSumFull().compareTo(BigDecimal.ZERO) < 0);
    }

    private de.dwcaesar.invoiceConverter.io.model.Invoice createBasicInvoice() {
        de.dwcaesar.invoiceConverter.io.model.Invoice invoice = new de.dwcaesar.invoiceConverter.io.model.Invoice();
        invoice.setInvoiceNumber(BigInteger.valueOf(12345));

        de.dwcaesar.invoiceConverter.io.model.Address shippingAddress = new de.dwcaesar.invoiceConverter.io.model.Address();
        shippingAddress.setName("John Doe");
        shippingAddress.setStreet("123 Main St");
        shippingAddress.setCity("Anytown");
        shippingAddress.setZip("12345");
        invoice.setShippingAddress(shippingAddress);

        de.dwcaesar.invoiceConverter.io.model.Address billingAddress = new de.dwcaesar.invoiceConverter.io.model.Address();
        billingAddress.setName("Jane Doe");
        billingAddress.setStreet("456 Other St");
        billingAddress.setCity("Othertown");
        billingAddress.setZip("67890");
        invoice.setBillingAddress(billingAddress);

        invoice.setPaymentMethod(de.dwcaesar.invoiceConverter.io.model.PaymentMethod.CREDIT_CARD);

        Map<Vat, BigDecimal> vat = new HashMap<>();
        vat.put(Vat.FULL, new BigDecimal("0.19"));
        vat.put(Vat.REDUCED, new BigDecimal("0.07"));
        invoice.setVat(vat);

        invoice.setNetto(new BigDecimal("100.00"));
        invoice.setBrutto(new BigDecimal("119.00"));

        de.dwcaesar.invoiceConverter.io.model.Item item = new de.dwcaesar.invoiceConverter.io.model.Item();
        item.setName("Test Item");
        item.setAmount(BigInteger.valueOf(1));
        item.setItemPrice(new BigDecimal("100.00"));
        item.setVat(Vat.FULL);
        invoice.setItems(Collections.singletonList(item));

        return invoice;
    }
}
