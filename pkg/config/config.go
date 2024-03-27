package config

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
)

var _ = pp.Println

type Config struct {
	endpoint string `yaml:"endpoint"`
	version  string `yaml:"version"`
	log      *logrus.Logger
}

func (c *Config) Log() *logrus.Logger {
	return c.log
}

func (c *Config) GetEndpoint() string {
	return c.endpoint
}

func (c *Config) GetVersion() string {
	return c.version
}

func (c *Config) ToString() string {
	return fmt.Sprintf("Config{endpoint: %s, version: %s}", c.endpoint, c.version)
}

func New(log *logrus.Logger) (config *Config, err error) {
	config = &Config{
		endpoint: "ws://poc-websocket.poc-websocket.svc.cluster.local/ws",
		version:  "1.0.1",
		log:      log,
	}
	return config, nil
}
