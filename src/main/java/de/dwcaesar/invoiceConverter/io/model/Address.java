package de.dwcaesar.invoiceConverter.io.model;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class Address {
    private String name;
    private String street;
    private String zip;
    private String city;
}
