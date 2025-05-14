package de.dwcaesar.invoiceConverter.config;

import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Contact;
import io.swagger.v3.oas.models.info.Info;
import io.swagger.v3.oas.models.info.License;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class OpenApiConfig {

    @Bean
    public OpenAPI invoiceConverterOpenAPI() {
        return new OpenAPI()
                .info(new Info()
                        .title("Invoice Converter API")
                        .description("API for converting invoice data between different formats")
                        .version("v1.0.0")
                        .contact(new Contact()
                                .name("Invoice Converter Team")
                                .email("support@invoiceconverter.example.com")
                                .url("https://github.com/dwcaesar/invoice-converter"))
                        .license(new License()
                                .name("MIT License")
                                .url("https://opensource.org/licenses/MIT")));
    }
}