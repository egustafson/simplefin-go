package main

import (
	"os"
	"time"
)

type Flags struct {
	Verbose   bool
	OutPath   string
	AccessURL string
	StartDate time.Time
	EndDate   time.Time
	Pending   bool
	BalOnly   bool
}

var cmdFlags Flags

const ENV_ACCESS_URL = "SF_ACCESS" //nolint:stylecheck

func init() {

	accessURL := os.Getenv(ENV_ACCESS_URL)
	if accessURL != "" {
		cmdFlags.AccessURL = accessURL
	}

	rootCmd.PersistentFlags().BoolVarP(
		&cmdFlags.Verbose,
		"verbose",
		"v",
		false,
		"Enable verbose output",
	)
	rootCmd.PersistentFlags().StringVarP(
		&cmdFlags.OutPath,
		"outfile",
		"o",
		"",
		"Path to output file (JSON). If not specified, output to STDOUT",
	)
	rootCmd.PersistentFlags().StringVar(
		&cmdFlags.AccessURL,
		"access-url",
		cmdFlags.AccessURL,
		"SimpleFIN access URL (or set "+ENV_ACCESS_URL+" environment variable)",
	)
	rootCmd.PersistentFlags().TimeVar(
		&cmdFlags.StartDate,
		"start-date",
		time.Time{},
		[]string{time.RFC3339, "2006-01-02"},
		"Start date for transactions (optional)",
	)
	rootCmd.PersistentFlags().TimeVar(
		&cmdFlags.EndDate,
		"end-date",
		time.Time{},
		[]string{time.RFC3339, "2006-01-02"},
		"End date for transactions (optional)",
	)
	rootCmd.PersistentFlags().BoolVar(
		&cmdFlags.Pending,
		"pending",
		false,
		"Include pending transactions",
	)
	rootCmd.PersistentFlags().BoolVar(
		&cmdFlags.BalOnly,
		"balances-only",
		false,
		"Return only account balances, no transactions",
	)
}
