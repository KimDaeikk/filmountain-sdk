package cmd

import (
	"github.com/spf13/cobra"
)

var minerContractCmd = &cobra.Command{
	Use: "miner-contract",
}

func init() {
	rootCmd.AddCommand(minerContractCmd)
}
