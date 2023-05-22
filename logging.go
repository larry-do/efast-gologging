package gologging

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"strings"
	"time"
)

func ConfigLogging(console bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.With().
		CallerWithSkipFrameCount(2).
		Int("pid", os.Getpid()).
		Logger()

	if console {
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
		log.Logger = log.Output(output)
	}
}