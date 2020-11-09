package cache

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strconv"
	"testing"
	"time"
)

type CmdableMock struct {
	redis.Cmdable
	mock.Mock
}

func (cmd *CmdableMock) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	args := cmd.Called(ctx, key, value, expiration)
	return args.Get(0).(*redis.StatusCmd)
}

func (cmd *CmdableMock) Get(ctx context.Context, key string) *redis.StringCmd {
	args := cmd.Called(ctx, key)
	return args.Get(0).(*redis.StringCmd)
}

func TestService_Set(t *testing.T) {
	key := 50
	val := map[int]int{10: 5}
	valJson, err := json.Marshal(val)
	assert.NoError(t, err)
	ctx := context.Background()
	cmd := redis.NewStatusCmd(ctx)
	cmdableMock := new(CmdableMock)
	cmdableMock.On("Set", mock.Anything, strconv.Itoa(key), string(valJson), mock.Anything).Once().Return(cmd)
	rdb := New(cmdableMock)
	rdb.Set(ctx, key, val)
}

func TestService_Get(t *testing.T) {
	val := `{"1":1}`
	ctx := context.Background()
	cmd := redis.NewStringResult(val, nil)
	cmdableMock := new(CmdableMock)
	cmdableMock.On("Get", mock.Anything, "1").Once().Return(cmd)
	rdb := New(cmdableMock)
	res := rdb.Get(ctx, 1)
	assert.NotNil(t, res)
	assert.Equal(t, map[int]int{1: 1}, res)
}

func TestService_Get_Error(t *testing.T) {
	key := 1
	ctx := context.Background()
	expectedErr := errors.New("test-error")
	cmd := redis.NewStringResult("", expectedErr)
	cmdableMock := new(CmdableMock)
	cmdableMock.On("Get", mock.Anything, strconv.Itoa(key)).Once().Return(cmd)
	rdb := New(cmdableMock)
	valRes := rdb.Get(ctx, key)
	assert.Nil(t, valRes)
}
