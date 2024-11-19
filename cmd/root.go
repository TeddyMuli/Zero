package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "zero",
	Short: "zero is a command line application for perfoming basic mathematical calculations",
	Long: "zero is a command line application for perfoming basic mathematical calculations - addition, substraction, division and multiplication",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops an error occured while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
