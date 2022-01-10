package manager

import (
	"fmt"
	"github.com/edwardsuwirya/go_dating/util/logger"
	"github.com/spf13/viper"
)

const (
	AppName    = "datingapp"
	ConfigName = "config"
)

type InfraManager interface {
	Config() *viper.Viper
}

type infraManager struct {
}

func (i *infraManager) Config() *viper.Viper {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("/etc/%s", AppName))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", AppName))
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Log.Fatal().Msg("Config File Not Found")
		} else {
			logger.Log.Fatal().Msg("Config File Error")
		}
	}
	return viper.GetViper()
}

func NewInfra() InfraManager {
	logger.NewLogger()
	newInfra := &infraManager{}
	return newInfra
}
