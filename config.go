package gologging

import (
	"fmt"
	"github.com/rs/zerolog"
	"strings"
)

type Config struct {
	ConsoleLoggingEnabled bool
	FileLoggingEnabled    bool
	GlobalLevel           string
	EncodeLogsAsJson      bool
	FilePath              string
	MaxSize               int
	MaxBackups            int
	MaxAge                int
	Compress              bool
	TimeZone              string
}

func (config Config) getConvetedLogLevel() zerolog.Level {
	if strings.EqualFold(config.GlobalLevel, "disabled") {
		return zerolog.Disabled
	}
	if strings.EqualFold(config.GlobalLevel, "no") {
		return zerolog.NoLevel
	}
	if strings.EqualFold(config.GlobalLevel, "panic") {
		return zerolog.PanicLevel
	}
	if strings.EqualFold(config.GlobalLevel, "fatal") {
		return zerolog.FatalLevel
	}
	if strings.EqualFold(config.GlobalLevel, "error") {
		return zerolog.ErrorLevel
	}
	if strings.EqualFold(config.GlobalLevel, "warn") {
		return zerolog.WarnLevel
	}
	if strings.EqualFold(config.GlobalLevel, "info") {
		return zerolog.InfoLevel
	}
	if strings.EqualFold(config.GlobalLevel, "debug") {
		return zerolog.DebugLevel
	}
	if strings.EqualFold(config.GlobalLevel, "trace") {
		return zerolog.TraceLevel
	}
	panic(fmt.Sprintf("Logging level %s not found", config.GlobalLevel))
}
