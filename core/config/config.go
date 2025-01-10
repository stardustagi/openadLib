package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/stardustagi/openadLib/core/logger"
	"github.com/stardustagi/openadLib/core/mongo"
	"github.com/stardustagi/openadLib/core/mysql"
	"github.com/stardustagi/openadLib/core/redis"
	"github.com/stardustagi/openadLib/service/http_service"
)

type Config struct {
	Logger *logger.Config       `json:"logger"`
	MySql  *mysql.Config        `json:"mysql"`
	Mongo  *mongo.Config        `json:"mongo"`
	Redis  *redis.Config        `json:"redis"`
	Http   *http_service.Config `json:"http"`
}

func ParseConfig(fn string) (*Config, error) {
	config := &Config{}
	_, err := toml.DecodeFile(fn, &config)
	if err != nil {
		return nil, fmt.Errorf("ParseConfig: %s", err)
	}
	return config, nil
}
