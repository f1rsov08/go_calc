package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		statusCode int
	}{
		{"1+1", 2, http.StatusOK},
		{"3 -4", -1, http.StatusOK},
		{"2 * 3", 6, http.StatusOK},
		{"5/ 10", 0.5, http.StatusOK},
		{"(1 + 2)* 3", 9, http.StatusOK},
		{"2+2*2", 6, http.StatusOK},
		{"(1 + 2", 0, http.StatusUnprocessableEntity},
		{"1 /0", 0, http.StatusInternalServerError},
		{"abc", 0, http.StatusUnprocessableEntity},
		{"1 + (2 * (3 - 1))", 5, http.StatusOK},
		{"", 0, http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			reqBody := map[string]string{"expression": test.expression}
			body, _ := json.Marshal(reqBody)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			CalcHandler(w, req)

			resp := w.Result()

			if resp.StatusCode != test.statusCode {
				t.Errorf("expected status %d; got %d", test.statusCode, resp.StatusCode)
			}

			if resp.StatusCode == http.StatusOK {
				var response map[string]float64
				json.NewDecoder(resp.Body).Decode(&response)
				if response["result"] != test.expected {
					t.Errorf("expected result %f; got %f", test.expected, response["result"])
				}
			}
		})
	}
}
