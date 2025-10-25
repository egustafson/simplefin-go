package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use: "about",
	Run: doAbout,
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}

func doAbout(cmd *cobra.Command, args []string) {
	fmt.Println("---")
	fmt.Printf("version:    %s\n", GitSummary)
	fmt.Printf("build-date: %s\n", BuildDate)
	fmt.Println("...")
}
