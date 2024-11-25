package config

import (
	"encoding/json"
	"fmt"
	"github.com/stardustagi/openadLib/core/logger"
	"io/ioutil"
)

type Config struct {
	Logger *logger.Config `json:"logger"`
}

type GameApi struct {
	RechargePay string `json:"recharge_pay"`
	SignIn      string `json:"sign_in"`
	TimeOut     int    `json:"time_out"`
}

func ParseConfig(path string) (*Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ParseConfig: %s", err)
	}
	config := &Config{}
	err = json.Unmarshal(content, config)
	if err != nil {
		return nil, fmt.Errorf("ParseConfig: %s", err)
	}
	return config, nil
}
