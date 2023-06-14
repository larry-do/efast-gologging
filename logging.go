package gologging

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"strings"
	"time"
)

func LoadLogginConfig(configFile string) {
	log.Info().Msgf("Reading logging configuration...")
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error loading logging configuration file")
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error parsing logging configuration file")
	}

	ConfigLogging(config)
}

func ConfigLogging(config Config) {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.SetGlobalLevel(config.getConvetedLogLevel())
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.DateTime,
			NoColor:    config.EncodeLogsAsJson,
		}
		if !config.EncodeLogsAsJson {
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
		}
		output.PartsExclude = []string{
			//zerolog.TimestampFieldName,
		}
		writers = append(writers, output)
	}

	if config.FileLoggingEnabled {
		writers = append(writers, &lumberjack.Logger{
			Filename:   config.FilePath,
			MaxBackups: config.MaxBackups, // files
			MaxSize:    config.MaxSize,    // megabytes
			MaxAge:     config.MaxAge,     // days
			Compress:   config.Compress,
		})
	}

	multiWriters := io.MultiWriter(writers...)
	log.Logger = zerolog.New(multiWriters).Level(config.getConvetedLogLevel()).
		With().Timestamp().
		CallerWithSkipFrameCount(2).
		Int("pid", os.Getpid()).
		Logger()
}
