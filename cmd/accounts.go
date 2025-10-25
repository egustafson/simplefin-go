package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/egustafson/simplefin-go"
	"github.com/spf13/cobra"
)

var accountsCmd = &cobra.Command{
	Use:  "accounts",
	RunE: doAccounts,
}

func init() {
	rootCmd.AddCommand(accountsCmd)
}

func doAccounts(cmd *cobra.Command, args []string) error {

	accessURL := cmdFlags.AccessURL
	if accessURL == "" {
		return fmt.Errorf("access URL is required (set with --access-url or %s environment variable)", ENV_ACCESS_URL)
	}

	sf, err := simplefin.New(accessURL)
	if err != nil {
		return fmt.Errorf("failed to create SimpleFin client: %w", err)
	}
	ar, err := sf.Accounts()
	if err != nil {
		return fmt.Errorf("failed to get accounts: %w", err)
	}

	encResp, err := json.Marshal(ar)
	if err != nil {
		return fmt.Errorf("failed to marshal accounts response: %w", err)
	}

	outFile := os.Stdout
	if cmdFlags.OutPath != "" {
		outFile, err = os.OpenFile(cmdFlags.OutPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("failed to open output file: %w", err)
		}
		defer outFile.Close()
	}
	fmt.Fprintln(outFile, string(encResp))
	return nil
}
