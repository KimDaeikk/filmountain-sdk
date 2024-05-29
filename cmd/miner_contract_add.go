package cmd

import (
	"time"

	"github.com/KimDaeikk/filmountain-sdk/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <miner actor address>",
	Short: "Add miner actor address to the miner contract",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.StartSpinner("test")
		time.Sleep(1 * time.Second)
		arg := "start"
		utils.SuccessfSpinner("sucess to %s test", arg)

		utils.StartSpinner("test")
		time.Sleep(1 * time.Second)
		utils.ErrorfSpinner("fail to %s test", arg)
	},
}

func init() {
	minerContractCmd.AddCommand(addCmd)
}
