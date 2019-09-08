package config

import (
	"github.com/stevenroose/gonfig"
)

// Config ...
type Config struct {
	ConfigFile string `id:"config" desc:"path to config"`
	Logger     struct {
		Level  string `id:"level" default:"info" desc:"loggin level"`
		Format string `id:"format" default:"console"`
	} `id:"logger"`
	MQ struct {
		URL           string    `id:"url"`
		AutoAck       bool      `id:"auto_ack"`
		ReconnectTime *Duration `id:"reconnect_time" default:"5m"`
	} `id:"mq"`
}

// Load loads config file
func Load(filename string) (*Config, error) {
	var cfg Config
	err := gonfig.Load(&cfg, gonfig.Conf{
		ConfigFileVariable:  "config",
		FileDefaultFilename: filename,
		EnvPrefix:           "rmq_",
		FileDecoder:         gonfig.DecoderTryAll,
	})

	return &cfg, err
}
