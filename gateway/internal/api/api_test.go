package api

import (
	"encoding/json"
	"github.com/Glide28/go_hw_05/ledger"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func perform(handler http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w
}

func TestTransactionWorkflow(t *testing.T) {
	handler := NewHandler(ledger.New())
	budget := perform(handler, http.MethodPost, "/api/budgets", `{"category":"еда","limit":1000}`)
	if budget.Code != http.StatusCreated {
		t.Fatalf("budget status: %d, body: %s", budget.Code, budget.Body.String())
	}
	tx := perform(handler, http.MethodPost, "/api/transactions", `{"amount":450,"category":"еда","description":"ланч","date":"2026-10-10"}`)
	if tx.Code != http.StatusCreated {
		t.Fatalf("transaction status: %d, body: %s", tx.Code, tx.Body.String())
	}
	list := perform(handler, http.MethodGet, "/api/transactions", "")
	var items []TransactionResponse
	if err := json.Unmarshal(list.Body.Bytes(), &items); err != nil || len(items) != 1 {
		t.Fatalf("unexpected list: %s, error: %v", list.Body.String(), err)
	}
}

func TestBudgetExceeded(t *testing.T) {
	handler := NewHandler(ledger.New())
	perform(handler, http.MethodPost, "/api/budgets", `{"category":"еда","limit":100}`)
	response := perform(handler, http.MethodPost, "/api/transactions", `{"amount":101,"category":"еда","description":"ужин","date":"2026-10-10"}`)
	if response.Code != http.StatusConflict || !strings.Contains(response.Body.String(), "budget exceeded") {
		t.Fatalf("expected 409, got %d: %s", response.Code, response.Body.String())
	}
}

func TestValidationAndEmptyList(t *testing.T) {
	handler := NewHandler(ledger.New())
	bad := perform(handler, http.MethodPost, "/api/transactions", `{"amount":0,"category":"","date":"bad"}`)
	if bad.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", bad.Code)
	}
	empty := perform(handler, http.MethodGet, "/api/budgets", "")
	if empty.Code != http.StatusOK || strings.TrimSpace(empty.Body.String()) != "[]" {
		t.Fatalf("expected empty array, got %d: %s", empty.Code, empty.Body.String())
	}
	if got := empty.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Fatalf("unexpected content type: %q", got)
	}
}
