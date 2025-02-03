package de.dwcaesar.invoiceConverter.invoice;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Address {
    private String name;
    private String street;
    private String city;
    private String zipcode;
}
