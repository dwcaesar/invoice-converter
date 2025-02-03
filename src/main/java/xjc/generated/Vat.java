
package generated;

import jakarta.annotation.Generated;
import jakarta.xml.bind.annotation.XmlEnum;
import jakarta.xml.bind.annotation.XmlEnumValue;
import jakarta.xml.bind.annotation.XmlType;


/**
 * <p>Java-Klasse für vat.
 * 
 * <p>Das folgende Schemafragment gibt den erwarteten Content an, der in dieser Klasse enthalten ist.
 * <pre>
 * &lt;simpleType name="vat"&gt;
 *   &lt;restriction base="{http://www.w3.org/2001/XMLSchema}string"&gt;
 *     &lt;enumeration value="full"/&gt;
 *     &lt;enumeration value="reduced"/&gt;
 *     &lt;enumeration value="none"/&gt;
 *   &lt;/restriction&gt;
 * &lt;/simpleType&gt;
 * </pre>
 * 
 */
@XmlType(name = "vat")
@XmlEnum
@Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
public enum Vat {

    @XmlEnumValue("full")
    FULL("full"),
    @XmlEnumValue("reduced")
    REDUCED("reduced"),
    @XmlEnumValue("none")
    NONE("none");
    private final String value;

    Vat(String v) {
        value = v;
    }

    public String value() {
        return value;
    }

    public static Vat fromValue(String v) {
        for (Vat c: Vat.values()) {
            if (c.value.equals(v)) {
                return c;
            }
        }
        throw new IllegalArgumentException(v);
    }

}
