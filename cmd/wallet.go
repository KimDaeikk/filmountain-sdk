package cmd

import (
	"github.com/spf13/cobra"
)

// walletCmd represents the wallet command
var walletCmd = &cobra.Command{
	Use:   "wallet",
}

func init() {
	rootCmd.AddCommand(walletCmd)
}
