# Developer Guide for Invoice Converter

## Table of Contents
1. [Introduction](#introduction)
2. [Project Setup](#project-setup)
3. [Architecture Overview](#architecture-overview)
4. [Code Organization](#code-organization)
5. [Development Workflow](#development-workflow)
6. [Tools and Utilities](#tools-and-utilities)
7. [Testing](#testing)
8. [API Documentation](#api-documentation)
9. [Contributing](#contributing)

## Introduction

The Invoice Converter is a Spring Boot application that converts invoice data between different formats (XML to JSON). The project also includes tools for generating test data and measuring data quality. This guide provides information for developers who want to contribute to or use the project.

### Project Goals

As stated in the README, the main goal of this project is to explore and demonstrate the possible applications of AI in generating test data for software development. The project focuses on a real-life use case of invoice data conversion and aims to:

- Gain a better understanding of how AI can help engineers boost the software development process
- Improve developer experience
- Measure the quality of AI-generated test data

## Project Setup

### Prerequisites

- Java 17 or higher
- Gradle
- Go (for the tools)
- Git

### Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/dwcaesar/invoice-converter.git
   cd invoice-converter
   ```

2. Build the project:
   ```
   ./gradlew build
   ```

3. Run the application:
   ```
   ./gradlew bootRun
   ```

The application will start on port 8080 by default.

## Architecture Overview

The Invoice Converter follows a layered architecture:

1. **Presentation Layer**: Handles HTTP requests and responses through REST controllers.
2. **Service Layer**: Contains the business logic for converting invoices.
3. **Model Layer**: Defines the data structures for both input (XML) and output (JSON) formats.

### Key Components

- **InvoiceController**: REST controller that exposes the API endpoint for converting invoices.
- **InvoiceConverter**: Service that handles the conversion logic between different invoice formats.
- **Model Classes**: Data structures for representing invoices, addresses, items, etc.
- **Tools**: Separate utilities for generating test data and measuring data quality.

## Code Organization

The project is organized into the following packages:

- `de.dwcaesar.invoiceConverter`: Root package
  - `.config`: Configuration classes
  - `.invoice`: Invoice conversion logic and output model
    - `.model`: Output model classes
  - `.io`: Input/output handling
    - `.model`: Input model classes

The tools are organized in separate directories:

- `tools/dq_calc`: Data quality calculation tool
- `tools/generator_client`: Test data generation tool

## Development Workflow

### Feature Development

1. Create a new branch from `main` for your feature:
   ```
   git checkout -b feature/your-feature-name
   ```

2. Implement your changes, following the code style and architecture.

3. Write tests for your changes.

4. Run tests to ensure your changes don't break existing functionality:
   ```
   ./gradlew test
   ```

5. Submit a pull request for review.

### Code Style

The project follows standard Java code style conventions. Use the provided IDE configurations for consistent formatting.

## Tools and Utilities

### Data Quality Calculator (dq_calc)

The `dq_calc` tool is used to measure the quality of generated test data. It supports the following metrics:

- **Completeness**: Checks if all required fields are present and not empty.
- **Consistency**: Checks for intra-record consistency (e.g., brutto > netto) and format consistency.

#### Usage

```
cd tools/dq_calc
make run
```

The tool reads invoice XML files from the `assets` directory and calculates metrics based on the configuration in `assets/config.yaml`.

### Test Data Generator (generator_client)

The `generator_client` tool generates test invoice data using a language model (LLM). It sends requests to an LLM with a prompt that includes the invoice schema, and writes the responses as XML files.

#### Usage

```
cd tools/generator_client
make run
```

The tool reads configuration from `assets/config.yml`, which specifies:
- Target directory for generated files
- Maximum number of files to generate
- Prompt template
- Schema file path
- Model to use
- LLM URL

## Testing

The project uses JUnit for unit testing. Tests are located in the `src/test` directory, mirroring the structure of the main code.

### Running Tests for the Main Application

```
./gradlew test
```

### Test Data

Test data is available in the `src/test/resources` directory, including sample XML and JSON files.

### Testing the Tools

Both the data quality calculator (dq_calc) and test data generator (generator_client) tools have their own test suites written in Go. These tests ensure that the tools function correctly and produce expected results.

#### Testing the Data Quality Calculator (dq_calc)

The dq_calc tool has unit tests that verify the functionality of its metrics calculations, including completeness and consistency checks for invoices, addresses, and items.

To run the tests for the dq_calc tool:

```
cd tools/dq_calc
go test -v
```

The tests cover various scenarios:
- Testing invoice completeness (all required fields are present and not empty)
- Testing consistency between netto and brutto values
- Testing consistency between item prices and total netto value
- Testing address completeness
- Testing item completeness

#### Testing the Test Data Generator (generator_client)

The generator_client tool has unit tests that verify its functionality for building request bodies, querying the generator, reading configuration, and writing responses to files.

To run the tests for the generator_client tool:

```
cd tools/generator_client
go test -v
```

The tests cover:
- Building request bodies with the correct format
- Querying the generator with proper HTTP requests
- Reading configuration from YAML files
- Writing generated responses to files

These tests use Go's httptest package to mock HTTP requests and responses, allowing for testing without requiring an actual LLM service.

## API Documentation

The application includes comprehensive API documentation using Swagger/OpenAPI.

### Accessing the Documentation

When the application is running, you can access the API documentation in the following ways:

- **Swagger UI**: Navigate to [http://localhost:8080/swagger-ui.html](http://localhost:8080/swagger-ui.html) in your web browser.
- **OpenAPI Specification**: The raw OpenAPI specification in JSON format is available at [http://localhost:8080/api-docs](http://localhost:8080/api-docs).

## Contributing

Contributions to the project are welcome. Please follow these steps:

1. Check the `docs/tasks.md` file for a list of improvement tasks.
2. Choose a task or propose a new one.
3. Follow the development workflow described above.
4. Submit a pull request with your changes.

### Reporting Issues

If you find a bug or have a suggestion for improvement, please create an issue in the GitHub repository.
