/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// changeOwnerCmd represents the changeOwner command
var changeOwnerCmd = &cobra.Command{
	Use:   "change-owner",
	Short: "Propose owner to vault contract",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("changeOwner called")
	},
}

func init() {
	vaultCmd.AddCommand(changeOwnerCmd)
}
