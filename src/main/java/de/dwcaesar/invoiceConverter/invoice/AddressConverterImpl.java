package de.dwcaesar.invoiceConverter.invoice;

import de.dwcaesar.invoiceConverter.invoice.model.Address;
import org.springframework.stereotype.Service;

import java.util.Optional;

/**
 * Implementation of the AddressConverter interface
 */
@Service
public class AddressConverterImpl implements AddressConverter {

    @Override
    public Address convertShippingAddress(de.dwcaesar.invoiceConverter.io.model.Address shippingAddress) {
        return Address.builder()
                .name(shippingAddress.getName())
                .street(shippingAddress.getStreet())
                .city(shippingAddress.getCity())
                .zipcode(shippingAddress.getZip())
                .build();
    }

    @Override
    public Optional<Address> convertBillingAddress(de.dwcaesar.invoiceConverter.io.model.Address billingAddress) {
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
}