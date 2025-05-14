package de.dwcaesar.invoiceConverter.io;

import de.dwcaesar.invoiceConverter.invoice.model.Invoice;
import de.dwcaesar.invoiceConverter.invoice.InvoiceConverter;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import io.swagger.v3.oas.annotations.tags.Tag;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/")
@RequiredArgsConstructor
@Tag(name = "Invoice Converter", description = "API for converting invoice data between different formats")
public class InvoiceController {

    private final InvoiceConverter invoiceConverter;

    @Operation(
        summary = "Global error handler",
        description = "Handles all exceptions thrown by the API endpoints"
    )
    @ApiResponse(
        responseCode = "400",
        description = "Error occurred during processing",
        content = @Content(
            mediaType = "text/plain",
            schema = @Schema(type = "string")
        )
    )
    @ExceptionHandler
    public ResponseEntity<String> handleError(Exception ex) {
        return ResponseEntity.badRequest().body(ex.getMessage());
    }

    @Operation(
        summary = "Convert XML invoice to JSON format",
        description = "Takes an XML invoice as input and converts it to a JSON format with a different structure"
    )
    @ApiResponses(value = {
        @ApiResponse(
            responseCode = "200", 
            description = "Invoice successfully converted",
            content = @Content(
                mediaType = "application/json",
                schema = @Schema(implementation = Invoice.class)
            )
        ),
        @ApiResponse(
            responseCode = "400", 
            description = "Invalid input or conversion error",
            content = @Content(
                mediaType = "text/plain",
                schema = @Schema(type = "string")
            )
        )
    })
    @PostMapping(value = "/toJson", consumes = "application/xml", produces = "application/json")
    public Invoice toJson(
            @Parameter(
                description = "XML invoice to be converted", 
                required = true,
                content = @Content(
                    mediaType = "application/xml",
                    schema = @Schema(implementation = de.dwcaesar.invoiceConverter.io.model.Invoice.class)
                )
            )
            @RequestBody de.dwcaesar.invoiceConverter.io.model.Invoice invoice) {
        return invoiceConverter.convert(invoice);
    }

}
