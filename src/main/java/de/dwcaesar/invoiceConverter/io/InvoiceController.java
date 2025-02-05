package de.dwcaesar.invoiceConverter.io;

import de.dwcaesar.invoiceConverter.invoice.model.Invoice;
import de.dwcaesar.invoiceConverter.invoice.InvoiceConverter;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/")
@RequiredArgsConstructor
public class InvoiceController {

    private final InvoiceConverter invoiceConverter;

    @ExceptionHandler
    public ResponseEntity<String> handleError(Exception ex) {
        return ResponseEntity.badRequest().body(ex.getMessage());
    }

    @PostMapping(value = "/toJson", consumes = "application/xml", produces = "application/json")
    public Invoice toJson(@RequestBody de.dwcaesar.invoiceConverter.io.model.Invoice invoice) {
        return invoiceConverter.convert(invoice);
    }

}
