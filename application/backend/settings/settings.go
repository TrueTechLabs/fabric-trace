package settings

import (
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./settings")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
