package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var substractCmd = &cobra.Command{
	Use: "substract",
	Aliases: []string{"sub"},
	Args: cobra.ExactArgs(2),
	Short: "Substract 2 numbers",
	Long: "Carry out substraction operation on 2 numbers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Substraction of %s and %s = %s", args[0], args[1], Substract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(substractCmd)
}
