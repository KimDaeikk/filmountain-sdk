/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// borrowCmd represents the borrow command
var borrowCmd = &cobra.Command{
	Use:   "borrow",
	Short: "borrow FIL from filmountain pool",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("borrow called")
	},
}

func init() {
	vaultCmd.AddCommand(borrowCmd)
}
