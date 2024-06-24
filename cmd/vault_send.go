/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send FIL to miner worker",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("send called")
	},
}

func init() {
	vaultCmd.AddCommand(sendCmd)
}
