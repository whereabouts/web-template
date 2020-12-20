package configure

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
)

const (
	defaultPort       = 8080
	defaultEnv        = "prod"
	fieldNamePort     = "Port"
	fieldNameEnv      = "Env"
	defaultConfigPath = "./config/application.json"
)

type IConfig interface {
	Set()
}

type DefaultConfig struct {
	Port int64  `json:"port"`
	Env  string `json:"env"`
}

var dConfig = DefaultConfig{defaultPort, defaultEnv}

func Load(conf IConfig) (*DefaultConfig, error) {
	// first, find the path from the command line parameter or the default directory
	cmdParam := GetCmdParam("c", defaultConfigPath)
	data, err := ioutil.ReadFile(cmdParam.String())
	// if fail to read the default configure file, that will appoint some default values
	if err != nil {
		if cmdParam.String() == defaultConfigPath {
			return DefaultLoad(), nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	confV := reflect.ValueOf(conf).Elem()
	if port := confV.FieldByName(fieldNamePort); port.IsValid() && port.Int() > reflect.ValueOf(0).Int() {
		dConfig.Port = port.Int()
	}
	if env := confV.FieldByName(fieldNameEnv); env.IsValid() && env.String() != reflect.ValueOf("").String() {
		dConfig.Env = env.String()
	}
	return &dConfig, nil
}

func DefaultLoad() *DefaultConfig {
	return &dConfig
}
