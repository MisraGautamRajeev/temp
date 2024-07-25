package temp

import (
	"github.com/rs/zerolog"
	"math/rand"
	"time"
)

var logLevels = []string{"debug", "info", "warn", "error", "fatal", "panic"}
var logLevelMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
	"panic": zerolog.PanicLevel,
}

func InitLogger() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
}

func GenerateRandomLevel() {
	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(len(logLevels))
	newLogLevel := logLevelMap[logLevels[index]]
	zerolog.SetGlobalLevel(newLogLevel)
}
