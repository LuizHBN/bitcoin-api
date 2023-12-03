package test

import (
	"bitcoin-klever-api/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HealthCheckHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status": "healthy"}`
	if rr.Body.String() != expected {
		t.Errorf("HealthCheckHandler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetBitcoinData(t *testing.T) {

	req, err := http.NewRequest("GET", "/bitcoin/{address}", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"address": "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetBitcoinData)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetBitcoinData returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestGetBalance(t *testing.T) {
	req, err := http.NewRequest("GET", "/balance/{address}", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"address": "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetBalance)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetBalance returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSendHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/send", strings.NewReader(`{"address":"bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r","amount":2000000}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.SendHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("SendHandler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestGetTransactionData(t *testing.T) {

	req, err := http.NewRequest("GET", "/transaction/{tx}", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"tx": "dc70b48d955836f47d0c42df405f92b919fd2f2090e70ca7672139122545f959"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetTransactionData)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetTransactionData returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
