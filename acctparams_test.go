package simplefin

import (
	"testing"
	"time"
)

func TestToQueryString_AllFields(t *testing.T) {
	start := time.Unix(1609459200, 0) // 2021-01-01
	end := time.Unix(1612137600, 0)   // 2021-02-01
	ap := &AccountParams{
		StartDate:    start,
		EndDate:      end,
		Pending:      true,
		Accounts:     []string{"acc1", "acc2"},
		BalancesOnly: true,
	}
	q := ap.ToQueryString()
	if q != "?start-date=1609459200&end-date=1612137600&pending=1&account=acc1&account=acc2&balances-only=1" {
		t.Errorf("unexpected query string: %s", q)
	}
}

func TestToQueryString_Empty(t *testing.T) {
	ap := &AccountParams{}
	if ap.ToQueryString() != "" {
		t.Errorf("expected empty query string, got: %s", ap.ToQueryString())
	}
}

func TestWithStartDate(t *testing.T) {
	d := time.Now()
	ap := &AccountParams{}
	WithStartDate(d)(ap)
	if !ap.StartDate.Equal(d) {
		t.Error("StartDate not set correctly")
	}
}

func TestWithEndDate(t *testing.T) {
	d := time.Now()
	ap := &AccountParams{}
	WithEndDate(d)(ap)
	if !ap.EndDate.Equal(d) {
		t.Error("EndDate not set correctly")
	}
}

func TestWithPending(t *testing.T) {
	ap := &AccountParams{}
	WithPending(true)(ap)
	if !ap.Pending {
		t.Error("Pending not set to true")
	}
}

func TestWithAccounts(t *testing.T) {
	ap := &AccountParams{}
	WithAccounts("foo")(ap)
	if len(ap.Accounts) != 1 || ap.Accounts[0] != "foo" {
		t.Error("Accounts not set correctly")
	}
}

func TestWithBalancesOnly(t *testing.T) {
	ap := &AccountParams{}
	WithBalancesOnly(true)(ap)
	if !ap.BalancesOnly {
		t.Error("BalancesOnly not set to true")
	}
}

func TestProcessAccountOptions(t *testing.T) {
	d := time.Now()
	ap := processAccountOptions([]AccountOption{
		WithStartDate(d),
		WithPending(true),
		WithAccounts("bar"),
	})
	if !ap.StartDate.Equal(d) || !ap.Pending || len(ap.Accounts) != 1 || ap.Accounts[0] != "bar" {
		t.Error("processAccountOptions did not apply options correctly")
	}
}
