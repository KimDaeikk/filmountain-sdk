package cmd

import (
	"fmt"

	"github.com/KimDaeikk/filmountain-sdk/config"
	"github.com/KimDaeikk/filmountain-sdk/connectors"
	"github.com/KimDaeikk/filmountain-sdk/utils"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <miner actor address>",
	Short: "Add miner actor address to the miner vault",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.StartSpinner("add miner to vault...")

		// lotus-miner ID address 인수로 받기
		// lotus-miner info 명령어의 "Miner: "부분에서 알 수 있음
		minerAddr, err := address.NewFromString(args[0])
		if err != nil {
			utils.ErrorfSpinner("%s", err)
		}

		// Lotus rpc call을 위한 client 연결
		lapi, closer, err := connectors.ConnectLotusClient(&config.AppConf)
		if err != nil {
			utils.ErrorfSpinner("%s", err)
		}
		defer closer()

		// // hex to [20]byte
		// // miner를 위임할 filmountain vault contract 주소
		// ethAddr, err := ethtypes.ParseEthAddress(MinerContract.String())
		// if err != nil {
		// 	utils.ErrorfSpinner("%s", err)
		// }

		// // [20]byte to filecoin address
		// delegated, err := ethAddr.ToFilecoinAddress()
		// if err != nil {
		// 	utils.ErrorfSpinner("%s", err)
		// }

		// // filecoin address to FVM ID address
		// id, err := lapi.StateLookupID(cmd.Context(), delegated, types.EmptyTSK)
		// if err != nil {
		// 	utils.ErrorfSpinner("%s", err)
		// }

		// ID가 보유한, 특정 miner의 정보 객체
		mi, err := lapi.StateMinerInfo(cmd.Context(), minerAddr, types.EmptyTSK)
		if err != nil {
			utils.ErrorfSpinner("%s", err)
		}

		// TODO pendingOwnerAddress가 있을때 change-owner 명령의 동작 체크(안될 수도 있을까봐)
		fmt.Printf("%+v\n", mi.SectorSize)
		utils.SuccessSpinner("sucessfully add miner")
	},
}

func init() {
	vaultCmd.AddCommand(addCmd)
}
