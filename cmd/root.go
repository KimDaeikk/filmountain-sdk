package cmd

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/KimDaeikk/filmountain-sdk/config"
	"github.com/KimDaeikk/filmountain-sdk/constants"
	"github.com/KimDaeikk/filmountain-sdk/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "filmountain-sdk",
	Short: "this application is sdk for filmountain",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	utils.GenerateSpinner()
}

func initConfig() {
	// .env 파일에 CONFIG_DIR가 설정되어있다면
	// 해당 dir를 cfgDir로 설정하고
	if os.Getenv("CONFIG_DIR") != "" {
		config.CfgDir = os.Getenv("CONFIG_DIR")
	} else {
		// .env에 CONFIG_DIR가 설정되어있지않다면
		// ~, 즉 유저 홈 디렉터리를 가져와서
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		// 홈 디렉터리 하위에 /.filmountain 디렉터리를 cfgDir로 설정
		config.CfgDir = fmt.Sprintf("%s/.filmountain", home)
		// 모든 사용자에게 권한 부여하여 디렉터리 생성
		err = os.MkdirAll(config.CfgDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create config directory: %v", err)
		}
	}

	// viper의 디렉터리를 위의 cfgDir로 설정하고
	// config파일 확장자를 yaml로 지정
	viper.AddConfigPath(config.CfgDir)
	viper.SetConfigName("setting")
	viper.SetConfigType("yaml")

	utils.NewKeyStore(fmt.Sprintf("%s/keystore", config.CfgDir))

	if err := utils.NewKeyStoreLegacy(fmt.Sprintf("%s/keys.toml", config.CfgDir)); err != nil {
		log.Fatal(err)
	}

	if err := utils.NewVaultStore(fmt.Sprintf("%s/vault.toml", config.CfgDir)); err != nil {
		log.Fatal(err)
	}

	if err := utils.NewAccountsStore(fmt.Sprintf("%s/accounts.toml", config.CfgDir)); err != nil {
		log.Fatal(err)
	}

	if err := utils.NewBackupsStore(fmt.Sprintf("%s/backups.toml", config.CfgDir)); err != nil {
		log.Fatal(err)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// setting.yaml을 읽으려고 시도
	err := viper.ReadInConfig()
	if err != nil {
		// 만약 setting.yaml이 없다면
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 기본 설정값 지정
			setDefaultConfig("", &constants.DefaultAppConf)
			// 위의 기본 설정을 바탕으로 setting.yaml 파일 생성
			viper.SafeWriteConfig()
			// 기본 설정 setting.yaml을 읽으려고 시도
			err := viper.ReadInConfig()
			if err != nil {
				log.Fatalf("Failed to read default setting.yaml: %v", err)
			}
		} else {
			log.Fatalf("Failed to read setting.yaml: %v", err)
		}
	}

	// 읽고 처리하기
	// 이제 config.AppConfig.LotusNode.Address 같은 방식으로 설정 변수를 읽을 수 있음
	err = viper.Unmarshal(&config.AppConf)
	if err != nil {
		log.Fatalf("Failed to unmarshal config data to struct: %v", err)
	}

}

// viper가 기본적으로 구조체의 필드마다 값을 설정해야해서
// 구조체가 변경되어도 자동으로 재귀적으로 처리되도록 보조함수 선언
func setDefaultConfig(prefix string, s interface{}) {
	val := reflect.Indirect(reflect.ValueOf(s)) // 포인터를 실제 값으로 변환
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		structField := typ.Field(i)

		// yaml 태그 추출
		tag := structField.Tag.Get("mapstructure")
		if tag == "" || tag == "-" {
			continue // 태그가 없거나 대시(-)로 처리된 경우 무시
		}

		// yaml 태그에 콤마(,)가 있는 경우, 첫 번째 부분만 사용
		tagParts := strings.Split(tag, ",")
		fieldName := tagParts[0] // 실제 태그 이름
		fieldPath := prefix + fieldName

		switch field.Kind() {
		case reflect.Struct:
			// 재귀적으로 구조체 필드 처리, 경로에 현재 필드 이름 추가
			setDefaultConfig(fieldPath+".", field.Addr().Interface())
		default:
			// Viper 기본값 설정, 전체 경로 사용
			viper.Set(fieldPath, field.Interface())
		}
	}
}
