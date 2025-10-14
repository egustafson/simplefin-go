// Package simplefin implements the
// [SimpleFIN](https://www.simplefin.org/protocol.html) protocol.
package simplefin

import "github.com/shopspring/decimal"

type Error string

type Account struct {
	Org              Organization
	ID               string
	Name             string
	Currency         string
	AvailableBalance decimal.Decimal
	BalanceDate      int64
	Transactions     []Transaction
	Extra            any
}

type Organization struct {
	Domain  string // reauired: maybe
	SfinURL string // required: yes
	Name    string // required: maybe
	URL     string // required: no
	ID      string // required: no
}

type Transaction struct {
	ID           string
	Posted       int64 // unix epoc timestamp
	Amount       decimal.Decimal
	Description  string
	TransactedAt int64
	Pending      bool
	Extra        any
}

type Currency struct {
	Name string
	Abbr string
}
