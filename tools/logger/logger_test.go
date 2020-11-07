package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tid := "testTransaction"
	ctx := context.WithValue(context.Background(), transactionIdKey, tid)
	logger := New(ctx)
	assert.Equal(t, tid, logger.Data[transactionIdKey])
}

func TestNewNotStringTid(t *testing.T) {
	tid := 1243
	ctx := context.WithValue(context.Background(), transactionIdKey, tid)
	logger := New(ctx)
	assert.NotEqual(t, tid, logger.Data[transactionIdKey])
}

func TestNewWithoutTid(t *testing.T) {
	logger := New(context.Background())
	assert.Len(t, logger.Data, 1)
}
