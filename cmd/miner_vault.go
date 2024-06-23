package cmd

import (
	"github.com/spf13/cobra"
)

var minerVaultCmd = &cobra.Command{
	Use: "miner-vault",
}

func init() {
	rootCmd.AddCommand(minerVaultCmd)
}
