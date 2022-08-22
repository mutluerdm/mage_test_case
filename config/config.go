package config

import (
	"mage_test_case/mlog"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title string
	Mongo MongoConfig `toml:"mongo"`
	Api   APIConfig   `toml:"api"`
}

type MongoConfig struct {
	Url      string `toml:"url"`
	Database string `toml:"database"`
	MaxOpen  int64  `toml:"max_open"`
	MaxIdle  int64  `toml:"max_idle"`
}

type APIConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func LoadConfigs() Config {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		mlog.Fatalln(err)
	}
	return conf
}
