package gologging

import (
	"fmt"
	"github.com/rs/zerolog"
	"strings"
)

type Config struct {
	ConsoleLoggingEnabled bool   `yaml:"console"`
	FileLoggingEnabled    bool   `yaml:"file"`
	GlobalLevel           string `yaml:"level"`
	EncodeLogsAsJson      bool   `yaml:"json_format"`
	FilePath              string `yaml:"file_path"`
	MaxSize               int    `yaml:"max_size"`
	MaxBackups            int    `yaml:"max_backups"`
	MaxAge                int    `yaml:"max_age"`
	Compress              bool   `yaml:"compress"`
	TimeZone              string `yaml:"time_zone"`
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
