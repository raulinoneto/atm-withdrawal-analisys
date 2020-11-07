package logger

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const transactionIdKey = "transaction_id"

func New(ctx context.Context) *logrus.Entry {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	transactionId, ok := ctx.Value(transactionIdKey).(string)
	if !ok || len(transactionId) <= 0 {
		transactionId = uuid.New().String()
	}
	return logrus.WithContext(ctx).WithField(transactionIdKey, transactionId).WithTime(time.Now().Add(time.Second))
}
