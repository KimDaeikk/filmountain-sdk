/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// changeWorkerCmd represents the changeWorker command
var changeWorkerCmd = &cobra.Command{
	Use:   "change-worker",
	Short: "Propose to change miner worker address",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("changeWorker called")
	},
}

func init() {
	vaultCmd.AddCommand(changeWorkerCmd)
}
