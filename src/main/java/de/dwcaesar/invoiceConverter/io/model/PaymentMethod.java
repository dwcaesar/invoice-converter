package de.dwcaesar.invoiceConverter.io.model;

import com.fasterxml.jackson.annotation.JsonValue;
import lombok.Getter;

@Getter
public enum PaymentMethod {
    CREDIT_CARD("credit-card"),
    INVOICE("invoice"),
    SURNAME("surname");

    private final String value;

    PaymentMethod(String value) {
        this.value = value;
    }

    @JsonValue
    public String getValue() {
        return value;
    }
}
