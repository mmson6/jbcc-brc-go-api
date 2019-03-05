package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ctxKey string

const keyLogger ctxKey = "logger"

func ContextWithField(ctx context.Context, name string, value interface{}) context.Context {
	log := Current(ctx)
	newLog := log.WithField(name, value)
	return context.WithValue(ctx, keyLogger, newLog)
}

func ContextWithFields(ctx context.Context, fields logrus.Fields) context.Context {
	log := Current(ctx)
	newLog := log.WithFields(fields)
	return context.WithValue(ctx, keyLogger, newLog)
}

func Current(ctx context.Context) logrus.FieldLogger {
	var log logrus.FieldLogger
	var ok bool

	val := ctx.Value(keyLogger)
	if log, ok = val.(logrus.FieldLogger); !ok {
		log = New()
	}

	return log
}

func Set(ctx context.Context, log logrus.FieldLogger) context.Context {
	return context.WithValue(ctx, keyLogger, log)
}
