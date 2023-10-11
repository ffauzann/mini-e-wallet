package app

import (
	"log"

	"github.com/ffauzann/mini-e-wallet/internal/model"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Database Database
	Cache    Cache
	App      *model.AppConfig
}

func (c *Config) Setup() {
	c.readConfigFile()

	err := c.Server.Logger.init()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = c.Database.prepare()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = c.Cache.prepare()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (c *Config) readConfigFile() {
	viper.SetConfigName("config")                // Name of config file (without extension)
	viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./internal/app")        // Optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil { // Find and read the config file
		log.Fatalln(err)
		return
	}

	if err := viper.Unmarshal(c); err != nil { // Unmarshal to struct
		log.Fatalln(err)
		return
	}
}
