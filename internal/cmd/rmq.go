package main

import (
	"os"
	"time"

	"rmq/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Revision = "unknown"
	Version  = "unknown"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Debug().Msg("This message appears only when log level set to Debug")

	log.Info().
		Str("version", Version).
		Str("revision", Revision).
		Int("pid", os.Getpid()).
		Msg("started")

	c, err := config.Load("config.example.json")
	if err != nil {
		log.Panic().Err(err).Msg("can't load config")
	}
	log.Info().Str("level", c.Logger.Level).Msg("logger")
	log.Info().Dur("timeout", time.Duration(*c.MQ.ReconnectTime)).Msg("duration")
}
