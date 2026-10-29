package ledger

import (
	"errors"
	"testing"
	"time"
)

func TestLedgerBudgetLimit(t *testing.T) {
	l := New()
	if _, err := l.SetBudget(Budget{Category: "food", Limit: 1000}); err != nil {
		t.Fatal(err)
	}
	if _, err := l.AddTransaction(Transaction{Amount: 700, Category: "food", Date: time.Now()}); err != nil {
		t.Fatal(err)
	}
	_, err := l.AddTransaction(Transaction{Amount: 400, Category: "food", Date: time.Now()})
	if !errors.Is(err, ErrBudgetExceeded) {
		t.Fatalf("expected ErrBudgetExceeded, got %v", err)
	}
}

func TestListMethodsReturnEmptySlices(t *testing.T) {
	l := New()
	if got := l.ListTransactions(); got == nil || len(got) != 0 {
		t.Fatalf("expected empty transaction slice, got %#v", got)
	}
	if got := l.ListBudgets(); got == nil || len(got) != 0 {
		t.Fatalf("expected empty budget slice, got %#v", got)
	}
}
