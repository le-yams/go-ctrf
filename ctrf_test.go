package ctrf

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredProperties(t *testing.T) {
	// Arrange
	report := Report{
		Results: &Results{
			Tool: &Tool{
				Name: "my tool",
			},
			Summary: &Summary{
				Tests:   15,
				Passed:  5,
				Failed:  4,
				Pending: 3,
				Skipped: 2,
				Other:   1,
				Start:   42,
				Stop:    1337,
			},
			Tests: []*Test{
				{
					Name:     "test 1",
					Status:   TestPassed,
					Duration: 10,
				},
				{
					Name:     "test 2",
					Status:   TestFailed,
					Duration: 11,
				},
				{
					Name:     "test 3",
					Status:   TestPending,
					Duration: 12,
				},
				{
					Name:     "test 4",
					Status:   TestSkipped,
					Duration: 13,
				},
				{
					Name:     "test 5",
					Status:   TestOther,
					Duration: 14,
				},
			},
		},
	}

	expectedJson := `{
	"results": {
		"tool": {
			"name": "my tool"
		},
		"summary": {
			"tests": 15,
			"passed": 5,
			"failed": 4,
			"pending": 3,
			"skipped": 2,
			"other": 1,
			"start": 42,
			"stop": 1337
		},
		"tests": [
			{
				"name": "test 1",
				"status": "passed",
				"duration": 10
			},
			{
				"name": "test 2",
				"status": "failed",
				"duration": 11
			},
			{
				"name": "test 3",
				"status": "pending",
				"duration": 12
			},
			{
				"name": "test 4",
				"status": "skipped",
				"duration": 13
			},
			{
				"name": "test 5",
				"status": "other",
				"duration": 14
			}
		]
	}
}`

	// Act
	actualJson := prettyJSON(t, report)

	// Assert
	assert.Equal(t, expectedJson, actualJson)
}

func prettyJSON(t *testing.T, report Report) string {
	jsonBytes, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}

	var buffer bytes.Buffer
	err = json.Indent(&buffer, jsonBytes, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	return string(buffer.Bytes())
}

type validationTestCase struct {
	Name           string   `json:"name"          yaml:"name"`
	Report         string   `json:"report"        yaml:"report"`
	ExpectedErrors []string `json:"expected_errors" yaml:"expected_errors"`
}

func TestValidation(t *testing.T) {
	data, err := os.ReadFile("validation-test-cases.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var testCases []validationTestCase
	err = yaml.Unmarshal(data, &testCases)
	if err != nil {
		t.Fatal(err)
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			report := Report{}
			err = json.Unmarshal([]byte(testCase.Report), &report)
			if err != nil {
				t.Fatal(err)
			}

			errs := report.Validate()
			if len(testCase.ExpectedErrors) == 0 {
				assert.Nil(t, errs)
			} else {
				assert.NotNil(t, errs)
				for _, expectedError := range testCase.ExpectedErrors {

					found := slices.ContainsFunc(errs, func(err error) bool {
						return err.Error() == expectedError
					})

					if !found {
						t.Error("Expected error not found:", expectedError)
					}
				}
			}
		})
	}
}
