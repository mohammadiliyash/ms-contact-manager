package config

import (
	"io/ioutil"

	log "github.com/miliyash/ms-contact-manager/logging"
	"gopkg.in/yaml.v2"
)

// Config for the micro-service
var Config *MSCMConfig

func Initialize() {
	log.Debug("calling config.Initialize()")
	confContent, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	conf := &MSCMConfig{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		panic(err)
	}
}

// SetConfig bypasses the usual config override and validation behavior.
// It is intended to be used by unit tests.
func SetConfig(conf *MSCMConfig) {
	Config = conf
}

type MSCMConfig struct {
	Environment string `yaml:"environment"`
	Key         string `yaml:"key"`
}

const AppName = "ms-contact-manager"
