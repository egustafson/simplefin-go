// Package simplefin implements the
// [SimpleFIN](https://www.simplefin.org/protocol.html) protocol.
package simplefin

// file simplefin.go implements the struct's and interfaces of the simplefin-go
// library.

import (
	_ "encoding/json"

	"github.com/shopspring/decimal"
)

// SimpleFin is the interface to interact with a SimpleFIN service.
type SimpleFin interface {
	//
	// Info queries the SimpleFIN server and returns the supported protocol
	// versions.
	Info() (InfoResp, error)
	//
	// Accounts is the primary method: it queries the SimpleFIN server and
	// returns a list of accounts with an optional list of transactions
	// attributed to each account.
	Accounts(options ...AccountOption) (AccountsResp, error)
	//
	// Claim() is provided by this package but is not a method of the SimpleFin
	// interface.
	//
	//  Claim(token string) (accessURL string, err error)
}

// InfoResp is the response to an Info() call to the server.  A list of
// supported SimpleFIN protocol versions is returned.
type InfoResp struct {
	Versions []string `json:"versions"`
}

// AccountsResp is the response to an Accounts() call to the server.  A list of
// accounts and, optionally, transactions for those accounts is returned.
type AccountsResp struct {
	Errors   []ReqError `json:"errors,omitempty"`
	Accounts []Account  `json:"accounts,omitempty"`
}

// ReqError is an error message as returned by a SimpleFIN server.
type ReqError string

// Account represents the details of a single account as returned by a SimpleFIN
// server.
type Account struct {
	Org              Organization    `json:"org"`
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	Currency         string          `json:"currency"`
	Balance          decimal.Decimal `json:"balance"`
	AvailableBalance decimal.Decimal `json:"available-balance"`
	BalanceDate      int64           `json:"balance-date"` // unix epoc timestamp
	Transactions     []Transaction   `json:"transactions,omitempty"`
	Extra            map[string]any  `json:"extra,omitempty"`
}

// Organization represents the details of a single organization as returned by a
// SimpleFIN server.
type Organization struct {
	Domain  string `json:"domain"`        // reauired: maybe
	SfinURL string `json:"sfin-url"`      // required: yes
	Name    string `json:"name"`          // required: maybe
	URL     string `json:"url,omitempty"` // required: no
	ID      string `json:"id,omitempty"`  // required: no
}

// Transaction represents a single account transaction as returned by a
// SimpleFIN server.
type Transaction struct {
	ID           string          `json:"id"`
	Posted       int64           `json:"posted"` // unix epoc timestamp
	Amount       decimal.Decimal `json:"amount"`
	Description  string          `json:"description"`
	TransactedAt int64           `json:"transacted_at"` // unix epoc timestamp
	Pending      bool            `json:"pending,omitempty"`
	Extra        map[string]any  `json:"extra,omitempty"`
}

// Currency represents a "Custom Currency" as queried from a currency URL in the
// Account structure returned by the Accounts() interface call.
type Currency struct {
	Name string `json:"name"`
	Abbr string `json:"abbr"`
}
