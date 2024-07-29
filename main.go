package temp

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

var myGlobalLogger = log.Logger

func InitLogger() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
}

func GenerateRandomLevel() {
	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(len(logLevels))
	newLogLevel := logLevelMap[logLevels[index]]
	zerolog.SetGlobalLevel(newLogLevel)
}

func GetPackageLogger() *zerolog.Logger {
	return &myGlobalLogger
}

func GenerateRandomGlobalLevel() {
	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(len(logLevels))
	ChangeGlobalLevel(logLevels[index])
}

func ChangeGlobalLevel(newLevel string) {
	level, ok := logLevelMap[newLevel]
	if !ok {
		return
	}
	myGlobalLogger = log.Logger.Level(level)
}

func ChangeGlobalLogger(newLevel string) {
	level, ok := logLevelMap[newLevel]
	if !ok {
		return
	}
	log.Logger = log.Logger.Level(level)
}
