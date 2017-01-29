package utils_test

import (
	"strings"
	"testing"

	"github.com/PennTex/commuter/cmd/utils"
)

func TestUtils_FormatDateTimeInput(t *testing.T) {
	cases := []struct {
		input                string
		expectError          bool
		expectErrorToContain string
		expectedResult       int64
	}{
		{
			input:                "",
			expectError:          true,
			expectErrorToContain: "Invalid date/time format",
		},
		{
			input:                "0400",
			expectError:          true,
			expectErrorToContain: "Invalid date/time format",
		},
		{
			input:                "0:0630",
			expectError:          true,
			expectErrorToContain: "Invalid date format",
		},
		{
			input:                "0401:0",
			expectError:          true,
			expectErrorToContain: "Invalid time format",
		},
		{
			input:                "0000:0000",
			expectError:          true,
			expectErrorToContain: "Invalid month supplied",
		},
		{
			input:                "0100:0000",
			expectError:          true,
			expectErrorToContain: "Invalid day supplied",
		},
		{
			input:                "0101:0000",
			expectError:          true,
			expectErrorToContain: "Invalid time supplied. Expected 24 hr time format.",
		},
		{
			input:                "0101:0000AM",
			expectError:          true,
			expectErrorToContain: "Invalid time supplied. Expected 12 hr time format.",
		},
	}

	for _, c := range cases {
		result, err := utils.FormatDateTimeInput(c.input)

		if c.expectError {
			if err == nil {
				t.Errorf("Expected %s to error", c.input)
			} else if !strings.Contains(err.Error(), c.expectErrorToContain) {
				t.Errorf("Expected error '%s' to contain '%s'", err.Error(), c.expectErrorToContain)
			}
		} else if err != nil {
			t.Errorf("Did not expected %s to error", c.input)
		} else {
			if result != c.expectedResult {
				t.Errorf("Expected result to be %d but it was %d", c.expectedResult, result)
			}
		}
	}
}
