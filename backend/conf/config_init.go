package conf

import (
	"log"

	"github.com/google/wire"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

var conf *Conf

func NewConf() *Conf {
	if conf != nil {
		return conf
	}
	initConfig()
	return conf
}

func GetConf() *Conf {
	if conf != nil {
		return conf
	}
	initConfig()
	return conf
}

func SetConf(data *Conf) {
	conf = data
}

func initConfig() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	configName := "config/config.yaml"

	err := config.LoadFiles(configName)
	if err != nil {
		panic(err)
	}

	data := Conf{}
	err = config.BindStruct("", &data)
	if err != nil {
		panic(err)
	}
	log.Printf("load config form :%s data:%+v\n", configName, data)
	conf = &data
}

// ProviderSet is controller providers.
var ProviderSet = wire.NewSet(
	NewConf,
)
