package de.dwcaesar.invoiceConverter.io.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;
import lombok.Getter;

@Getter
public enum Vat {
    FULL("full"),
    REDUCED("reduced"),
    NONE("none");
    private final String value;

    Vat(String value) {
        this.value = value;
    }

    @JsonValue
    public String getValue() {
        return value;
    }
}
