package manager

import (
	"context"
	"fmt"
	"github.com/edwardsuwirya/go_dating/util/logger"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
)

const (
	AppName    = "datingapp"
	ConfigName = "config"
)

type InfraManager interface {
	Config() *viper.Viper
	SqlDb() *pgx.Conn
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

func (i *infraManager) SqlDb() *pgx.Conn {
	dbName := i.Config().GetString("datingapp.db.name")
	dbHost := i.Config().GetString("datingapp.db.host")
	dbPort := i.Config().GetString("datingapp.db.port")
	dbUser := i.Config().GetString("datingapp.db.user")
	dbPassword := i.Config().GetString("datingapp.db.password")

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := pgx.Connect(context.Background(), dataSourceName)
	if err != nil {
		logger.Log.Fatal().Msg("DB failed to start")
	}
	return conn
}

func NewInfra() InfraManager {
	logger.NewLogger()
	newInfra := &infraManager{}
	return newInfra
}
