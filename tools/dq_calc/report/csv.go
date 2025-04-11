package report

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

/*
	Representation of a report of the test results.

	'TestID' - Unique identifier for each test run
	'Timestamp' - when this test method was executed
	'DataSourceName' - reference to the tested data point: file name + dom path
	'TestMethodUsed' - Identifier of the test method
	'ExpectedResult' - Result that was expected by the test method
	'ActualResult' - Result that was present in the data to test
	'Passed' - boolean, true if passed
	'Comment' - optional, additional info for the test method

	generated using ollama qwen2.5-coder, reviewed and refined manually
*/

type Report struct {
	TestEntries []TestEntry
	file        *os.File
	writer      *csv.Writer
}

type TestEntry struct {
	TestID         string
	Timestamp      time.Time
	DataSourceName string
	TestMethodUsed string
	ExpectedResult string
	ActualResult   string
	Passed         bool
	Comment        string
}

func NewReport(path string) (*Report, error) {
	fileName := filepath.Join(path, fmt.Sprintf("report_%d.csv", time.Now().Unix()))
	file, err := os.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf("error creating report file: %v", err)
	}
	writer := csv.NewWriter(file)

	header := []string{"TestID", "Timestamp", "DataSourceName", "TestMethodUsed", "ExpectedResult", "ActualResult", "Passed", "Comment"}
	if err := writer.Write(header); err != nil {
		return nil, fmt.Errorf("error writing report header: %v", err)
	}
	writer.Flush()

	return &Report{
		TestEntries: []TestEntry{},
		file:        file,
		writer:      writer,
	}, nil
}

func (report *Report) NewEntry(dataSourceName string, testMethodUsed string, expectedResult string) {
	entry := TestEntry{
		TestID:         fmt.Sprintf("test_%d", len(report.TestEntries)+1),
		Timestamp:      time.Now(),
		DataSourceName: dataSourceName,
		TestMethodUsed: testMethodUsed,
		ExpectedResult: expectedResult,
		Passed:         false,
		Comment:        "",
	}
	report.TestEntries = append(report.TestEntries, entry)
}

func (report *Report) CloseEntry(actualResult string, passed bool, comment string) {
	if len(report.TestEntries) == 0 {
		return
	}
	entry := report.TestEntries[len(report.TestEntries)-1]
	entry.ActualResult = actualResult
	entry.Passed = passed
	entry.Comment = comment
	report.TestEntries[len(report.TestEntries)-1] = entry

	row := []string{
		entry.TestID,
		entry.Timestamp.Format(time.RFC3339),
		entry.DataSourceName,
		entry.TestMethodUsed,
		entry.ExpectedResult,
		entry.ActualResult,
		fmt.Sprintf("%t", passed),
		entry.Comment,
	}
	if err := report.writer.Write(row); err != nil {
		fmt.Printf("error writing report row: %v", err)
	}

	report.writer.Flush()
}

func (report *Report) CloseReport() error {
	if err := report.writer.Error(); err != nil {
		return fmt.Errorf("error writing report file: %v", err)
	}
	if err := report.file.Close(); err != nil {
		return fmt.Errorf("error closing report file: %v", err)
	}
	return nil
}
