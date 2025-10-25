package main

import (
	"fmt"
	"os"

	"github.com/egustafson/simplefin-go"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:  "info",
	RunE: doInfo,
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

func doInfo(cmd *cobra.Command, args []string) error {

	accessURL := cmdFlags.AccessURL
	if accessURL == "" {
		return fmt.Errorf("access URL is required (set with --access-url or %s environment variable)", ENV_ACCESS_URL)
	}

	sf, err := simplefin.New(accessURL)
	if err != nil {
		return fmt.Errorf("failed to create SimpleFin client: %w", err)
	}
	ir, err := sf.Info()
	if err != nil {
		return err
	}

	ver := "[ "
	for i, v := range ir.Versions {
		if i != 0 {
			ver += ", "
		}
		ver += "\"" + v + "\""
	}
	ver += " ]"

	outFile := os.Stdout
	if cmdFlags.OutPath != "" {
		outFile, err = os.OpenFile(cmdFlags.OutPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("failed to open output file: %w", err)
		}
		defer outFile.Close()
	}
	fmt.Fprintln(outFile, "---")
	fmt.Fprintf(outFile, "versions: %s\n", ver)
	fmt.Fprintln(outFile, "...")
	return nil
}
