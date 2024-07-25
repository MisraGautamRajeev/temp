package temp

import "github.com/rs/zerolog"

func InitLogger() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
}
