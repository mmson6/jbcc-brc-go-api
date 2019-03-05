package logger

import (
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func New() logrus.FieldLogger {
	log := logrus.New()

	// Format
	log.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
		TimestampFormat:  time.RFC3339Nano,
		PrettyPrint:      false,
	})

	// Level
	level := logrus.InfoLevel
	if levelEnv, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if parsed, err := logrus.ParseLevel(levelEnv); err == nil {
			level = parsed
		}
	} else if debugEnv, ok := os.LookupEnv("DEBUG"); ok {
		debug, _ := strconv.ParseBool(debugEnv)
		if debug {
			level = logrus.DebugLevel
		}
	}
	log.SetLevel(level)

	return log
}
