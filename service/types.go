package service

import "context"

type (
	BillCount map[int]int

	Cache interface {
		Get(ctx context.Context, key string) map[int]int
		Set(ctx context.Context, key string, val interface{})
	}
)