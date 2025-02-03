
package generated;

import jakarta.annotation.Generated;
import jakarta.xml.bind.annotation.XmlEnum;
import jakarta.xml.bind.annotation.XmlEnumValue;
import jakarta.xml.bind.annotation.XmlType;


/**
 * <p>Java-Klasse für method.
 * 
 * <p>Das folgende Schemafragment gibt den erwarteten Content an, der in dieser Klasse enthalten ist.
 * <pre>
 * &lt;simpleType name="method"&gt;
 *   &lt;restriction base="{http://www.w3.org/2001/XMLSchema}string"&gt;
 *     &lt;enumeration value="credit card"/&gt;
 *     &lt;enumeration value="invoice"/&gt;
 *     &lt;enumeration value="surname"/&gt;
 *   &lt;/restriction&gt;
 * &lt;/simpleType&gt;
 * </pre>
 * 
 */
@XmlType(name = "method")
@XmlEnum
@Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
public enum Method {

    @XmlEnumValue("credit card")
    CREDIT_CARD("credit card"),
    @XmlEnumValue("invoice")
    INVOICE("invoice"),
    @XmlEnumValue("surname")
    SURNAME("surname");
    private final String value;

    Method(String v) {
        value = v;
    }

    public String value() {
        return value;
    }

    public static Method fromValue(String v) {
        for (Method c: Method.values()) {
            if (c.value.equals(v)) {
                return c;
            }
        }
        throw new IllegalArgumentException(v);
    }

}
