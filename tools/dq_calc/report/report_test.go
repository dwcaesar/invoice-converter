package report

import (
	"bufio"
	"daquam/assert"
	"fmt"
	"os"
	"testing"
	"time"
)

// generated using ollama qwen2.5-coder, reviewed and refined manually

func TestNewReport(t *testing.T) {
	t.Run("Should create a new report file with the correct header", func(t *testing.T) {
		path := t.TempDir()
		report, err := NewReport(path)
		if err != nil {
			t.Errorf("Error creating report: %v", err)
		}
		defer func(report *Report) {
			_ = report.CloseReport()
		}(report)

		expectedHeader := "TestID,Timestamp,DataSourceName,TestMethodUsed,ActualResult,Passed,Comment"
		actualHeader, err := report.ReadLine(1)
		if err != nil {
			t.Errorf("Error reading header: %v", err)
		}

		assert.Equal(t, actualHeader, expectedHeader)
	})
}

func TestNewEntry(t *testing.T) {
	t.Run("Should add a new entry to the test entries", func(t *testing.T) {
		path := t.TempDir()
		report, err := NewReport(path)
		if err != nil {
			t.Errorf("Error creating report: %v", err)
		}
		defer func(report *Report) {
			_ = report.CloseReport()
		}(report)

		dataSourceName := "test_data.csv"
		testMethodUsed := "TestExample"

		report.NewEntry(dataSourceName, testMethodUsed)

		assert.Equal(t, len(report.TestEntries), 1)
		entry := report.TestEntries[0]
		assert.Equal(t, entry.DataSourceName, dataSourceName)
		assert.Equal(t, entry.TestMethodUsed, testMethodUsed)
	})
}

func TestCloseEntry(t *testing.T) {
	t.Run("Should update the last entry with actual result and passed status", func(t *testing.T) {
		path := t.TempDir()
		report, err := NewReport(path)
		if err != nil {
			t.Errorf("Error creating report: %v", err)
		}
		defer func(report *Report) {
			_ = report.CloseReport()
		}(report)

		dataSourceName := "test_data.csv"
		testMethodUsed := "TestExample"

		report.NewEntry(dataSourceName, testMethodUsed)

		actualResult := "actualResult"
		comment := "passed as expected"

		report.CloseEntry(actualResult, true, comment)
		expectedLine := fmt.Sprintf("test_1,%s,test_data.csv,TestExample,actualResult,true,passed as expected", time.Now().Format(time.RFC3339))

		assert.Equal(t, len(report.TestEntries), 1)
		entry := report.TestEntries[0]
		assert.Equal(t, entry.ActualResult, actualResult)
		assert.Equal(t, entry.Passed, true)
		assert.Equal(t, entry.Comment, comment)

		actualLine, err := report.ReadLine(2)

		assert.Equal(t, actualLine, expectedLine)
	})
}

func (report *Report) ReadLine(lineNumber int) (string, error) {
	if lineNumber <= 0 {
		return "", fmt.Errorf("line number must be greater than zero")
	}
	file, err := os.Open(report.file.Name())
	if err != nil {
		return "", fmt.Errorf("error opening report for reading: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)

	for i := 1; scanner.Scan(); i++ {
		if i == lineNumber {
			return scanner.Text(), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading report file: %v", err)
	}

	return "", nil
}
