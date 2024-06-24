package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/KimDaeikk/filmountain-sdk/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/AlecAivazis/survey/v2"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// 키 존재여부 체크
		panicIfKeyExists(utils.OwnerKey)
		panicIfKeyExists(utils.OperatorKey)

		// 개인키 AES 대칭키 암호화용 비밀번호 생성
		ownerPassphrase, envSet := os.LookupEnv("FILMOUNTAIN_OWNER_PASSPHRASE")
		if !envSet {
			prompt := &survey.Password{
				Message: "Please type a passphrase to encrypt your owner private key",
			}
			survey.AskOne(prompt, &ownerPassphrase)
			var confirmPassphrase string
			confirmPrompt := &survey.Password{
				Message: "Confirm passphrase",
			}
			survey.AskOne(confirmPrompt, &confirmPassphrase)
			if ownerPassphrase != confirmPassphrase {
				log.Fatal("Aborting. Passphrase confirmation did not match.")
			}
		}

		utils.StartSpinner("add miner to vault...")
		ks := utils.KeyStore()

		// 이더리움 공개키 개인키 생성
		owner, err := ks.NewAccount(ownerPassphrase)
		if err != nil {
			log.Fatal(err)
		}

		// 환경변수가 없는 상태면 빈문자열 ""이 passphrase로 사용
		operatorPassphrase := os.Getenv("FILMOUNTAIN_OPERATOR_PASSPHRASE")
		operator, err := ks.NewAccount(operatorPassphrase)
		if err != nil {
			log.Fatal(err)
		}

		// 이더리움 공개키 저장
		as := utils.AccountsStore()

		as.Set(string(utils.OwnerKey), owner.Address.String())
		as.Set(string(utils.OperatorKey), operator.Address.String())

		if err := viper.WriteConfig(); err != nil {
			log.Fatal(err)
		}

		ownerAddr, ownerDelAddr, err := as.GetAddrs(string(utils.OwnerKey))
		if err != nil {
			log.Fatal(err)
		}
		operatorAddr, operatorDelAddr, err := as.GetAddrs(string(utils.OperatorKey))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Owner address: %s (ETH), %s (FIL)\n", ownerAddr, ownerDelAddr)
		log.Printf("Operator address: %s (ETH), %s (FIL)\n", operatorAddr, operatorDelAddr)
		log.Println()
		log.Println("Please make sure to fund your Owner Address with FIL before creating a Vault")

		bs := utils.BackupsStore()
		bs.Invalidate()
	},
}

func panicIfKeyExists(key utils.KeyType) {
	as := utils.AccountsStore()
	if as == nil {
		log.Fatal("AccountsStore is not initialized")
	}
	_, _, err := as.GetAddrs(string(key))
	if err == nil {
		log.Fatal("owner account already created")
	} else {
		var e *utils.ErrKeyNotFound
		if !errors.As(err, &e) {
			log.Fatal(err)
		}
	}
}

func init() {
	walletCmd.AddCommand(newCmd)
}
