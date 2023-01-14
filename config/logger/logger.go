package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {

	//TODO change using flags
	//zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	//Test messages
	log.Error().Msg("Error message")
	log.Warn().Msg("Warning message")
	log.Info().Msg("Info message")
	log.Debug().Msg("Debug message")
	log.Trace().Msg("Trace message")
}
