package config

import (
	"flag"

	gcfg "gopkg.in/gcfg.v1"
)

type (
	Configuration struct {
		SlackAPI SlackAPIConfig
	}

	SlackAPIConfig struct {
		Token string
	}
)

var (
	cfg Configuration

	configFilepath string
)

func init() {
	flag.StringVar(&configFilepath, "config", "", "path to config file")
}

func Load() error {
	if !flag.Parsed() {
		flag.Parse()
	}

	return gcfg.ReadFileInto(&cfg, configFilepath)
}

func Get() Configuration {
	return cfg
}
