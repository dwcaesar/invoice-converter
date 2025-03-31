# invoice-converter

## Background

This repo serves the purpose of exploring and demonstrating the possible applications of AI in the use case of generating test data, as one of the most work and time intensive task in the software development process. The team aims at hand-on approach and adapts a real-life use case to consider different aspects of using AI as a supportive tool for test data generation.

## Approach

We want to discover possible pitfalls, limitations and emphasise advantages of AI's usage when generating test data. To illustrate pros and cons the results are compared with outcomes from some already established approaches and techniques e.g. manual creation of test data and random test data generation.

We are not aiming for the completeness of our endeavour and approach.

### Goals

**The main goal** is to gain a **better understanding** how AI can help engineers in the particular use case **boost the software development process** and improve **dev experience** and ultimately to have measurable results on it. The journey is the goal as well.

Our goals **are not** comparing different LLMs, AI tuning and estimation of the various prompt techniques.

### The use case
We choose a real-life use case, which deals with lots of very sensitive data, e.g. PID, health data. Due to the limited resource capacity only a tight but crucial part of the real-life use case is going to be taken into consideration.


### Measurements
When measuring the quality of the generated test data we consider following common metrics [1]:

1. Completeness
2. Unambiguous
3. Correctness
4. Timeliness
5. Accuracy
6. Consistency
7. Redundancy
8. Relevance
9. Uniformity
10. Reliability

Furthermore, we take into account following aspects of AI as a test data genration tool
* Availability
* Integrability
* Pricing
* Licensing costs

### Moving Goal Posts

As this project proves to cover a larger area than anticipated, we break it down into smaller tasks.
That helps us to focus our efforts on the overarching goal without getting lost in the deep.

For that we decided on the following two-step-process.
We focus on improving the data quality to meet a defined threshold.
When this threshold is reached, we switch focus on introducing new metrics or refining existing metrics.
After the metrics have been refined, we switch back on improving the data quality.

Since this back and forth can create a moving goal post, it is necessary to adhere to this process instead of pursuing both aspects simultaneously.

To start the process some groundwork is required first:
* Set up a method to generate and persist testdata automatically,
* Measure data quality of generated data and persist its results.

Testdata can be generated and persisted using [generator client](tools/generator_client/main.go).
The data quality can be measured using [dq calc](tools/dq_calc/main.go).
As of [#7](https://github.com/dwcaesar/invoice-converter/pull/7) dq_calc supports the following metrics:
* completeness
  * Address with Name, Street, ZIP, and City set
  * Item with Name, Amount, Item price, and VAT set
* consistency
  * Total sum netto is consistent with item prices and their respective amount
  * Total sum brutto is greater than total sum netto

Given these metrics, we start the development process with improving data quality until 70 % of all generated data are "complete" and "consistent".

## Conclusion
TBD

## Sources

[1]: <https://quality.nfdi4ing.de/en/latest/general_quality/general_quality.html> "General Data Quality Metrics"
[2]: <https://github.com/sdv-dev/SDV?tab=readme-ov-file> "SDV"
[3]: <https://mostly.ai> "Mostly AI"
[4]: <https://www.softwaretestingmaterial.com/test-data-generator-tools/> "8 Best Test Data Generator Tools"
