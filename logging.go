package gologging

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

func ConfigLogging(config Config) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.SetGlobalLevel(config.getConvetedLogLevel())
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			NoColor:    false,
		}
		output.FormatLevel = func(i interface{}) string {
			switch i.(string) {
			case "info":
				return fmt.Sprintf("%5s:", color.LightGreen.Render(strings.ToUpper(i.(string))))
			case "debug":
				return fmt.Sprintf("%5s:", color.LightGreen.Render(strings.ToUpper(i.(string))))
			case "error":
				return fmt.Sprintf("%5s:", color.LightRed.Render(strings.ToUpper(i.(string))))
			case "fatal":
				return fmt.Sprintf("%5s:", color.LightRed.Render(strings.ToUpper(i.(string))))
			case "panic":
				return fmt.Sprintf("%5s:", color.LightRed.Render(strings.ToUpper(i.(string))))
			default:
				return fmt.Sprintf("%5s:", color.LightYellow.Render(strings.ToUpper(i.(string))))
			}
		}
		output.FormatMessage = func(i interface{}) string {
			if i == nil {
				return ""
			}
			return fmt.Sprintf("%s |", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s=", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		output.FormatTimestamp = func(i interface{}) string {
			unixTime, _ := i.(json.Number).Int64()
			return fmt.Sprintf("%s", time.UnixMilli(unixTime).Format("2006-01-02 15:04:05.000"))
		}
		output.PartsExclude = []string{
			//zerolog.TimestampFieldName,
		}
		writers = append(writers, output)
	}

	if config.FileLoggingEnabled {
		writers = append(writers, &lumberjack.Logger{
			Filename:   path.Join(config.Directory, config.Filename),
			MaxBackups: config.MaxBackups, // files
			MaxSize:    config.MaxSize,    // megabytes
			MaxAge:     config.MaxAge,     // days
			Compress:   config.Compress,
		})
	}

	multiWriters := io.MultiWriter(writers...)
	log.Logger = zerolog.New(multiWriters).With().
		CallerWithSkipFrameCount(2).
		Int("pid", os.Getpid()).
		Logger()
}
