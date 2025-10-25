// Package main implements the CLI entrypoint for the sf application.
package main

import (
	"os"

	"github.com/spf13/cobra"
)

// populated by linker at build time
var (
	// GitSummary = git describe --tags --dirty --always
	GitSummary string = "v0.0.0-dirty"

	// BuildDate = date -u +%Y-%m-%dT%H:%M:%SZ
	BuildDate string = "1970-01-01T00:00:00Z"
)

var (
	rootCmd = &cobra.Command{
		Use: "sf <sub-command>",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		PersistentPreRunE: initAppHook,
	}
)

// flags initialized in flags.go

func initAppHook(cmd *cobra.Command, args []string) (err error) {
	// app initialization
	return nil
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		// cobra will print an error to stdout/(?)err
		os.Exit(1)
	}
}
