package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestService_ProcessAmount(t *testing.T) {
	svc := New()
	for name, test := range tCases {
		bills := svc.ProcessAmount(context.Background(), test.payload)
		assert.NotNil(t, bills, "shouldn't be nil", name)
		assert.Equal(t, test.expected, bills)
	}
}
