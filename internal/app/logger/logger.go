package logger

import (
	"log"
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLog() {
	f, err := os.OpenFile("/tmp/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	consoleWriter := zerolog.New(f).Level(zerolog.DebugLevel)
	multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout)
	Logger = zerolog.New(multi).With().Timestamp().Logger()
}
