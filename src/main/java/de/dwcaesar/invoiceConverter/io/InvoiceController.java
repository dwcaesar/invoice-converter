package de.dwcaesar.invoiceConverter.io;

import de.dwcaesar.invoiceConverter.invoice.Invoice;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/")
public class InvoiceController {

    @ExceptionHandler
    public ResponseEntity<String> handleError(Exception ex) {
        // TODO - Implement
        return ResponseEntity.badRequest().body(ex.getMessage());
    }

    @PostMapping(value = "/toJson", consumes = "application/xml", produces = "application/json")
    public Invoice toJson(@RequestBody Invoice invoice) {
        // TODO - Implement
        return Invoice.builder().build();
    }

}
