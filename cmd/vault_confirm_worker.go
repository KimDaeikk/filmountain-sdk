package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// confirmWorkerCmd represents the confirmWorker command
var confirmWorkerCmd = &cobra.Command{
	Use:   "confirm-worker",
	Short: "Accept to change miner worker address",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("confirmWorker called")
	},
}

func init() {
	vaultCmd.AddCommand(confirmWorkerCmd)
}
