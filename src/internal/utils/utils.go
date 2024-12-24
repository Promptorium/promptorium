package utils

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

/*
 * ---------------- Logging Utils ----------------
 */

func GetLogger() zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.PanicLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.FormattedLevels = getFormattedLevels()
	return log.Output(getConsoleWriter()).With().Caller().Logger()
}

func getFormattedLevels() map[zerolog.Level]string {
	return map[zerolog.Level]string{
		zerolog.TraceLevel: "TRACE",
		zerolog.DebugLevel: "DEBUG",
		zerolog.InfoLevel:  " INFO",
		zerolog.WarnLevel:  " WARN",
		zerolog.ErrorLevel: "ERROR",
		zerolog.FatalLevel: "FATAL",
		zerolog.PanicLevel: "PANIC",
	}
}

func getConsoleWriter() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05", FormatCaller: callerFormatter}
}

func callerFormatter(i interface{}) string {
	var MAX_LEN = 40
	var parts = []string{}
	var lineNumber = strings.Split(i.(string), ":")[1]
	var path = strings.Split(i.(string), ":")[0]

	for _, part := range strings.Split(path, "/") {
		if part == "cmd" || len(parts) > 0 || part == "internal" {
			parts = append(parts, part)
		}
	}
	outputLen := len(strings.Join(parts, "/")) + len(lineNumber)

	if outputLen > MAX_LEN {
		// Progressively substitute parts of the path with ellipsis until the output length is less than MAX_LEN (or the end of the path is reached)
		for i := 1; i < len(parts); i++ {
			if outputLen > MAX_LEN {
				outputLen -= len(parts[i])
				outputLen += 2
				parts[i] = ".."
			}
		}
	}
	result := strings.Join(parts, "/")
	if len(result) < MAX_LEN {
		result += strings.Repeat(" ", MAX_LEN-len(result)-len(lineNumber))
	}
	result = result + " :" + lineNumber + " |"
	return result
}

// Cached Data

type CachedData[T any] struct {
	name           string
	contentChannel chan T
	content        T
	isCached       bool
	refresh        func(chan T)
}

func NewCachedData[T any](refresh func(chan T), name string) CachedData[T] {
	newCachedData := CachedData[T]{
		name:           name,
		isCached:       false,
		refresh:        refresh,
		contentChannel: make(chan T),
	}

	go newCachedData.refresh(newCachedData.contentChannel)

	return newCachedData
}

func (c *CachedData[T]) GetContent() T {
	if !c.isCached {
		c.content = <-c.contentChannel
		c.isCached = true
	}

	return c.content
}
