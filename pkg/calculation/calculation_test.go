package calculation

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		shouldFail bool
	}{
		{"1+1", 2, false},
		{"3 -4", -1, false},
		{"2 * 3", 6, false},
		{"5/ 10", 0.5, false},
		{"(1 + 2)* 3", 9, false},
		{"2+2*2", 6, false},
		{"(1 + 2", 0, true},
		{"1 /0", 0, true},
		{"abc", 0, true},
		{"1 + (2 * (3 - 1))", 5, false},
		{"", 0, true},
	}

	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			result, err := Calc(test.expression)
			if test.shouldFail {
				if err == nil {
					t.Errorf("Expected error for expression: %s, but got none", test.expression)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect error for expression: %s, but got: %v", test.expression, err)
				} else if result != test.expected {
					t.Errorf("For expression: %s, expected: %f, but got: %f", test.expression, test.expected, result)
				}
			}
		})
	}
}
