package settings

import (
	"strings"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./settings")

	// 自动读取环境变量，支持通过环境变量覆盖配置
	viper.SetEnvPrefix("FABRIC_TRACE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
