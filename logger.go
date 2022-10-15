package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type Logger struct{}

func init() {
	// tmp
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func (logger *Logger) LogError(err error, format string, args ...interface{}) {
	log.Error().Stack().Err(err).Msgf(format, args...)
}

func (logger *Logger) LogFatal(err error, format string, args ...interface{}) {
	log.Fatal().Stack().Err(err).Msgf(format, args...)
}

func (logger *Logger) LogInfo(format string, args ...interface{}) {
	log.Info().Msgf(format, args...)
}

func (logger *Logger) LogWarn(format string, args ...interface{}) {
	log.Warn().Msgf(format, args...)
}

func (logger *Logger) Printf(format string, values ...interface{}) {
	log.Printf(format, values...)
}
