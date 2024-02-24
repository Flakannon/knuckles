package env

import (
	"github.com/spf13/viper"
)

func GetAppVersion() string {
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	return viper.GetString("VERSION")
}
