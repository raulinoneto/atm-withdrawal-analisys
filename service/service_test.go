package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CacheMock struct {
	mock.Mock
}

func (c *CacheMock) Get(ctx context.Context, key int) map[int]int {
	args := c.Called(ctx, key)
	if res := args.Get(0); res != nil {
		return res.(map[int]int)
	}
	return nil
}

func (c *CacheMock) Set(context.Context, int, map[int]int) {}

type tCase struct {
	payload     int
	expected    map[int]int
	expectedErr error
}

var tCases = map[string]tCase{
	"Success Amount 1": {
		payload:  1,
		expected: map[int]int{1: 1},
	},
	"Success Amount 5": {
		payload:  5,
		expected: map[int]int{5: 1},
	},
	"Success Amount 10": {
		payload:  10,
		expected: map[int]int{10: 1},
	},
	"Success Amount 50": {
		payload:  50,
		expected: map[int]int{50: 1},
	},
	"Success Amount 6": {
		payload:  6,
		expected: map[int]int{5: 1, 1: 1},
	},
	"Success Amount 11": {
		payload:  11,
		expected: map[int]int{10: 1, 1: 1},
	},
	"Success Amount 51": {
		payload:  51,
		expected: map[int]int{50: 1, 1: 1},
	},
	"Success Amount 16": {
		payload:  16,
		expected: map[int]int{10: 1, 5: 1, 1: 1},
	},
	"Success Amount 66": {
		payload:  66,
		expected: map[int]int{50: 1, 10: 1, 5: 1, 1: 1},
	},
	"Success Amount 87": {
		payload:     87,
		expected:    map[int]int{50: 1, 10: 3, 5: 1, 1: 2},
		expectedErr: nil,
	},
	"Success Amount 93": {
		payload:     93,
		expected:    map[int]int{50: 1, 10: 4, 1: 3},
		expectedErr: nil,
	},
	"Success Amount 993": {
		payload:     993,
		expected:    map[int]int{50: 19, 10: 4, 1: 3},
		expectedErr: nil,
	},
}

var cache = new(CacheMock)

func TestService_ProcessAmount(t *testing.T) {
	for name, tCase := range tCases {
		cache.On("Get", mock.Anything, tCase.payload).Once().Return(nil)
		svc := New(cache)
		res := svc.ProcessAmount(context.Background(), tCase.payload)
		assert.NotNil(t, res, "shouldn't be nil", name)
		assert.Equal(t, tCase.expected, res.Coins)
		assert.Equal(t, float64(tCase.payload), res.Amount)
	}
}

func TestService_ProcessAmountWithCache(t *testing.T) {
	for name, tCase := range tCases {
		cache.On("Get", mock.Anything, tCase.payload).Once().Return(tCase.expected)
		svc := New(cache)
		res := svc.ProcessAmount(context.Background(), tCase.payload)
		assert.NotNil(t, res, "shouldn't be nil", name)
		assert.Equal(t, tCase.expected, res.Coins)
		assert.Equal(t, float64(tCase.payload), res.Amount)
	}
}

func ExampleService_ProcessAmount() {
	amount := 1987
	cache.On("Get", mock.Anything, amount).Once().Return(nil)
	svc := New(cache)
	res := svc.ProcessAmount(context.Background(), amount)
	fmt.Printf("%+v",res)
	// Output: &{Amount:1987 Coins:map[1:2 5:1 10:3 50:39]}
}

func BenchmarkService_ProcessAmount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		cache.On("Get", mock.Anything, n).Once().Return(nil)
		svc := New(cache)
		_ = svc.ProcessAmount(context.Background(), n)
	}
}
