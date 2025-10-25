package simplefin

import (
	"fmt"
	"time"
)

// file accountparams.go implements the HTTP parameters available on the
// /accounts request

type AccountOption func(*AccountParams)

type AccountParams struct {
	StartDate    time.Time
	EndDate      time.Time
	Pending      bool
	Accounts     []string // list of account-id's
	BalancesOnly bool
}

func (ap *AccountParams) ToQueryString() string {
	q := "?"

	if !ap.StartDate.IsZero() {
		q += fmt.Sprintf("start-date=%d&", ap.StartDate.Unix())
	}

	if !ap.EndDate.IsZero() {
		q += fmt.Sprintf("end-date=%d&", ap.EndDate.Unix())
	}

	if ap.Pending {
		q += "pending=1&"
	}

	for _, accountID := range ap.Accounts {
		q += fmt.Sprintf("account=%s&", accountID)
	}

	if ap.BalancesOnly {
		q += "balances-only=1&"
	}

	return q[:len(q)-1] // remove trailing &
}

func WithStartDate(d time.Time) AccountOption {
	return func(ap *AccountParams) {
		ap.StartDate = d
	}
}

func WithEndDate(d time.Time) AccountOption {
	return func(ap *AccountParams) {
		ap.EndDate = d
	}
}

func WithPending(f bool) AccountOption {
	return func(ap *AccountParams) {
		ap.Pending = f
	}
}

func WithAccounts(accountID string) AccountOption {
	return func(ap *AccountParams) {
		if ap.Accounts == nil {
			ap.Accounts = make([]string, 0, 1)
		}
		ap.Accounts = append(ap.Accounts, accountID)
	}
}

func WithBalancesOnly(f bool) AccountOption {
	return func(ap *AccountParams) {
		ap.BalancesOnly = f
	}
}

func processAccountOptions(options []AccountOption) *AccountParams {
	ap := new(AccountParams)
	for _, opt := range options {
		opt(ap)
	}
	return ap
}
