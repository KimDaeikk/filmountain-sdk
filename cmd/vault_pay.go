package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// payCmd represents the pay command
var payCmd = &cobra.Command{
	Use:   "pay",
	Short: "Repay borrow FIL",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pay called")
	},
}

func init() {
	vaultCmd.AddCommand(payCmd)
}
