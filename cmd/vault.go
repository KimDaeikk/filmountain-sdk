package cmd

import (
	"github.com/spf13/cobra"
)

var vaultCmd = &cobra.Command{
	Use: "vault",
}

func init() {
	rootCmd.AddCommand(vaultCmd)
}
