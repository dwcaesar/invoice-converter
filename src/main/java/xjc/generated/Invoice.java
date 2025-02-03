
package generated;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;
import jakarta.annotation.Generated;
import jakarta.xml.bind.annotation.XmlAccessType;
import jakarta.xml.bind.annotation.XmlAccessorType;
import jakarta.xml.bind.annotation.XmlSchemaType;
import jakarta.xml.bind.annotation.XmlType;


/**
 * <p>Java-Klasse für invoice complex type.
 * 
 * <p>Das folgende Schemafragment gibt den erwarteten Content an, der in dieser Klasse enthalten ist.
 * 
 * <pre>
 * &lt;complexType name="invoice"&gt;
 *   &lt;complexContent&gt;
 *     &lt;restriction base="{http://www.w3.org/2001/XMLSchema}anyType"&gt;
 *       &lt;choice&gt;
 *         &lt;element name="invoiceNumber" type="{http://www.w3.org/2001/XMLSchema}integer"/&gt;
 *         &lt;element name="billingAddress" type="{}address" minOccurs="0"/&gt;
 *         &lt;element name="shippingAddress" type="{}address"/&gt;
 *         &lt;element name="paymentMethod" type="{}method"/&gt;
 *         &lt;element name="item" type="{}item" maxOccurs="unbounded"/&gt;
 *         &lt;element name="netto" type="{http://www.w3.org/2001/XMLSchema}decimal"/&gt;
 *         &lt;element name="brutto" type="{http://www.w3.org/2001/XMLSchema}decimal"/&gt;
 *       &lt;/choice&gt;
 *     &lt;/restriction&gt;
 *   &lt;/complexContent&gt;
 * &lt;/complexType&gt;
 * </pre>
 * 
 * 
 */
@XmlAccessorType(XmlAccessType.FIELD)
@XmlType(name = "invoice", propOrder = {
    "invoiceNumber",
    "billingAddress",
    "shippingAddress",
    "paymentMethod",
    "item",
    "netto",
    "brutto"
})
@Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
public class Invoice {

    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    protected BigInteger invoiceNumber;
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    protected Address billingAddress;
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    protected Address shippingAddress;
    @XmlSchemaType(name = "string")
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    protected Method paymentMethod;
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    protected List<Item> item;
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    protected BigDecimal netto;
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    protected BigDecimal brutto;

    /**
     * Ruft den Wert der invoiceNumber-Eigenschaft ab.
     * 
     * @return
     *     possible object is
     *     {@link BigInteger }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public BigInteger getInvoiceNumber() {
        return invoiceNumber;
    }

    /**
     * Legt den Wert der invoiceNumber-Eigenschaft fest.
     * 
     * @param value
     *     allowed object is
     *     {@link BigInteger }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public void setInvoiceNumber(BigInteger value) {
        this.invoiceNumber = value;
    }

    /**
     * Ruft den Wert der billingAddress-Eigenschaft ab.
     * 
     * @return
     *     possible object is
     *     {@link Address }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public Address getBillingAddress() {
        return billingAddress;
    }

    /**
     * Legt den Wert der billingAddress-Eigenschaft fest.
     * 
     * @param value
     *     allowed object is
     *     {@link Address }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public void setBillingAddress(Address value) {
        this.billingAddress = value;
    }

    /**
     * Ruft den Wert der shippingAddress-Eigenschaft ab.
     * 
     * @return
     *     possible object is
     *     {@link Address }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public Address getShippingAddress() {
        return shippingAddress;
    }

    /**
     * Legt den Wert der shippingAddress-Eigenschaft fest.
     * 
     * @param value
     *     allowed object is
     *     {@link Address }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public void setShippingAddress(Address value) {
        this.shippingAddress = value;
    }

    /**
     * Ruft den Wert der paymentMethod-Eigenschaft ab.
     * 
     * @return
     *     possible object is
     *     {@link Method }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public Method getPaymentMethod() {
        return paymentMethod;
    }

    /**
     * Legt den Wert der paymentMethod-Eigenschaft fest.
     * 
     * @param value
     *     allowed object is
     *     {@link Method }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public void setPaymentMethod(Method value) {
        this.paymentMethod = value;
    }

    /**
     * Gets the value of the item property.
     * 
     * <p>
     * This accessor method returns a reference to the live list,
     * not a snapshot. Therefore any modification you make to the
     * returned list will be present inside the Jakarta XML Binding object.
     * This is why there is not a <CODE>set</CODE> method for the item property.
     * 
     * <p>
     * For example, to add a new item, do as follows:
     * <pre>
     *    getItem().add(newItem);
     * </pre>
     * 
     * 
     * <p>
     * Objects of the following type(s) are allowed in the list
     * {@link Item }
     * 
     * 
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public List<Item> getItem() {
        if (item == null) {
            item = new ArrayList<>();
        }
        return this.item;
    }

    /**
     * Ruft den Wert der netto-Eigenschaft ab.
     * 
     * @return
     *     possible object is
     *     {@link BigDecimal }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public BigDecimal getNetto() {
        return netto;
    }

    /**
     * Legt den Wert der netto-Eigenschaft fest.
     * 
     * @param value
     *     allowed object is
     *     {@link BigDecimal }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public void setNetto(BigDecimal value) {
        this.netto = value;
    }

    /**
     * Ruft den Wert der brutto-Eigenschaft ab.
     * 
     * @return
     *     possible object is
     *     {@link BigDecimal }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public BigDecimal getBrutto() {
        return brutto;
    }

    /**
     * Legt den Wert der brutto-Eigenschaft fest.
     * 
     * @param value
     *     allowed object is
     *     {@link BigDecimal }
     *     
     */
    @Generated(value = "com.sun.tools.xjc.Driver", comments = "JAXB RI v3.0.2", date = "2025-02-03T12:35:35+01:00")
    public void setBrutto(BigDecimal value) {
        this.brutto = value;
    }

}
