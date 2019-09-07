package main

import (
	"os"

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
}
