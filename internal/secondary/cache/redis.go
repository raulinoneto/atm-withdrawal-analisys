package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/raulinoneto/atm-withdrawal-analisys/tools/logger"
)

type Service struct {
	svc redis.Cmdable
}

func New(svc redis.Cmdable) *Service {
	return &Service{svc}
}

func (rs *Service) Set(ctx context.Context, key int, value map[int]int) {
	log := logger.New(ctx)
	log.Info("Set on redis")
	if err := rs.svc.Set(ctx, strconv.Itoa(key), value, 24*time.Hour).Err(); err != nil {
		log.Error("Error on set on Redis")
		return
	}
	log.Info("Set on Success")
	return
}

func (rs *Service) Get(ctx context.Context, key int) map[int]int {
	log := logger.New(ctx)
	log.Info("Get from redis")
	val := rs.svc.Get(ctx, strconv.Itoa(key))
	result, err := val.Bytes()
	if err != nil && err != redis.Nil {
		log.Error("Error on get on Redis")
		return nil
	}
	log.WithField("result", result)
	log.Info("Got from redis")
	res := make(map[int]int)
	if err := json.Unmarshal(result, &res); err != nil {
		log.Error("Error unmarsal response")
		return nil
	}
	return res
}
