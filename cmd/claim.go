package main

import (
	"fmt"

	"github.com/egustafson/simplefin-go"
	"github.com/spf13/cobra"
)

var (
	claimCmd = &cobra.Command{
		Use:   "claim",
		Short: "Claim a SimpleFIN access URL using a claim token",
		RunE:  doClaim,
	}

	tokenFlag string
)

func init() {
	rootCmd.AddCommand(claimCmd)
	claimCmd.Flags().StringVarP(
		&tokenFlag,
		"token", // long flag
		"t",     // short flag
		"",      // default value
		"a claim token issued by SimpleFIN",
	)
}

func doClaim(cmd *cobra.Command, args []string) error {

	if tokenFlag == "" {
		return fmt.Errorf("claim token must be provided")
	}

	accessURL, err := simplefin.Claim(tokenFlag)
	if err != nil {
		return err
	}

	fmt.Println("---")
	fmt.Println("access-url:", accessURL)
	fmt.Println("...")

	return nil
}
