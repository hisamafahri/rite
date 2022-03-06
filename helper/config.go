package helper

import (
	"github.com/spf13/viper"
)

type ConfigStruct struct {
	Groups map[string]interface{} `yaml:"groups"`
	Users  map[string]interface{} `yaml:"users"`
}

func LoadConfig() (config ConfigStruct, err error) {
	// Read the config file
	viper.AddConfigPath(".rite")
	viper.SetConfigName("rite.config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
