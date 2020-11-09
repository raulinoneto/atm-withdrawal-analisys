package logger

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const transactionIdKey = "transaction_id"

type Logger struct {
	*logrus.Entry
}

func New(ctx context.Context) *Logger {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	transactionId, ok := ctx.Value(transactionIdKey).(string)
	if !ok || len(transactionId) <= 0 {
		transactionId = uuid.New().String()
		ctx = context.WithValue(ctx, transactionIdKey, transactionId)
	}
	return &Logger{
		Entry: logrus.WithContext(ctx).
			WithField(transactionIdKey, transactionId).
			WithTime(time.Now().Add(time.Second)),
	}
}

func (l *Logger) WithField(key string, value interface{}) {
	l.Entry = l.Entry.WithField(key, value)
}
