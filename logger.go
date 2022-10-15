package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	envManaber "github.com/punk-link/environment-variable-manager"
)

type Logger struct{}

func init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logLevel := getLogLevel()
	switch strings.ToLower(logLevel) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}

	host, err := os.Hostname()
	if err != nil {
		log.Logger = log.With().Str("host", "unknown").Logger()
	} else {
		log.Logger = log.With().Str("host", host).Logger()
	}

	log.Logger = log.With().Stack().Caller().Logger()
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

func getLogLevel() string {
	isExist, level := envManaber.TryGetEnvironmentVariable("LOG_LEVEL")
	if !isExist {
		return "Error"
	}

	return level
}
