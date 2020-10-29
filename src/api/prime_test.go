package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPreLargestPrimeLookUp(t *testing.T) {
	testSet := []struct {
		target   int
		expected int
	}{
		{-1, -1}, // invalid
		{0, -1},  // invalid
		{1, -1},  // invalid
		{2, 2},
		{5, 5},
		{25, 23},
		{55, 53},
		{100, 97},
	}

	for _, testCase := range testSet {
		prime := preLargestPrimeLookUp(testCase.target)
		if prime != testCase.expected {
			t.Errorf("Func returned wrong result: got %v want %v",
				prime, testCase.expected)
		}
	}
}

func TestPrimeHandler(t *testing.T) {
	data, _ := json.Marshal(map[string]int{"target": 55})

	req, err := http.NewRequest("POST", "/api/prime", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(PrimeHandler)
	handler.ServeHTTP(rec, req)

	var res responseMessage
	_ = json.NewDecoder(rec.Body).Decode(&res)

	// Check the status code
	if res.Status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v with message %s want %v",
			res.Status, res.Message, http.StatusOK)
	}

	// Check the response body
	expected := 53
	if res.Prime != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Prime, expected)
	}
}

func TestPrimeHandlerInvalidDataType(t *testing.T) {
	// Receive invalid data type of Target number
	data, _ := json.Marshal(map[string]string{"target": "fdafad"})

	req, err := http.NewRequest("POST", "/api/prime", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(PrimeHandler)
	handler.ServeHTTP(rec, req)

	// Check the response
	var res responseMessage
	_ = json.NewDecoder(rec.Body).Decode(&res)

	// Check the status code
	if res.Status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v with message %s want %v",
			res.Status, res.Message, http.StatusBadRequest)
	}
}

func TestPrimeHandlerInvalidDataValue(t *testing.T) {
	// Receive invalid data value of Target number
	data, _ := json.Marshal(map[string]int{"target": 1})

	req, err := http.NewRequest("POST", "/api/prime", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(PrimeHandler)
	handler.ServeHTTP(rec, req)

	// Check the response
	var res responseMessage
	_ = json.NewDecoder(rec.Body).Decode(&res)

	// Check the status code
	if res.Status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v with message %s want %v",
			res.Status, res.Message, http.StatusBadRequest)
	}
}
