package main

import (
	"daquam/assert"
	"testing"
)

func TestEmailOk(t *testing.T) {
	var testString = "2@test.de"

	value := validateEmail(testString)
	assert.Equal(t, value, true)
}

func TestEmailFaild(t *testing.T) {
	var testString = "2/@test.de"

	value := validateEmail(testString)
	assert.Equal(t, value, false)

}

func TestFalseCompleteness(t *testing.T) {
	var emps = []Employee{
		{Id: "1",
			Birthdate: "s",
			Hiredate:  "b",
			Email:     "",
		},
		{Id: "",
			Birthdate: "x",
			Hiredate:  "a",
			Email:     "b",
		}, {Id: "1",
			Birthdate: "",
			Hiredate:  "s",
			Email:     "f",
		}, {Id: "a",
			Birthdate: " ",
			Hiredate:  "s",
			Email:     "f",
		},
	}
	var fields = []string{"X", "Hiredate"}
	for _, emp := range emps {
		data, err := convertEmployeeToMap(emp)
		if err != nil {
			t.Fail()
		}
		value := validateCompleteness(fields, data)
		assert.Equal(t, value, false)
	}
}

func TestTrueCompleteness(t *testing.T) {
	var emp = Employee{Id: "1",
		Birthdate: "s",
		Hiredate:  "b",
		Email:     "a",
	}
	data, err := convertEmployeeToMap(emp)
	if err != nil {
		t.Fail()
	}
	value := validateCompleteness([]string{}, data)
	assert.Equal(t, value, true)
}

func TestCalcMetricEmptyInput(t *testing.T) {
	value := calculateMetric(Consistency, nil, nil)
	assert.Equal(t, value, 0)
}

func TestCalcMetricOutput(t *testing.T) {
	value := calculateMetric(Consistency, []string{}, []map[string]string{})
	assert.Equal(t, value, 0)
}
