package container

import (
	"context"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/raulinoneto/atm-withdrawal-analisys/internal/httpserver"
	"github.com/raulinoneto/atm-withdrawal-analisys/internal/primary/v1.withdrawalhttp"
	"github.com/raulinoneto/atm-withdrawal-analisys/internal/secondary/cache"
	"github.com/raulinoneto/atm-withdrawal-analisys/pkg/domain/v1.withdrawal"
	"github.com/raulinoneto/atm-withdrawal-analisys/tools/logger"
)

type Container struct {
	server      *httpserver.Server
	httpadapter *withdrawalhttp.Adapter
	service     withdrawalhttp.WithdrawalService
	cache       withdrawal.Cacher
}

func (c *Container) GetServer(routes []httpserver.Route, middlewares []httpserver.MiddlewareFunc) *httpserver.Server {
	if c.server == nil {
		port := os.Getenv("PORT")
		if len(port) <=0 {
			port = "8080"
		}
		host := os.Getenv("HOST")
		if len(host) <=0 {
			host = "0.0.0.0"
		}
		c.server = httpserver.New(&httpserver.Options{
			Middlewares: middlewares,
			Routes:      routes,
			Port:        port,
			Host:        host,
			Logger: logger.New(context.Background()),
		})
	}
	return c.server
}

func (c *Container) GetHttpAdapter() *withdrawalhttp.Adapter {
	if c.httpadapter == nil {
		c.httpadapter = withdrawalhttp.New(c.GetService())
	}
	return c.httpadapter
}

func (c *Container) GetService() withdrawalhttp.WithdrawalService {
	if c.service == nil {
		c.service = withdrawal.New(c.GetCache())
	}
	return c.service
}

func (c *Container) GetCache() withdrawal.Cacher {
	if c.cache == nil {
		host := os.Getenv("REDIS_HOST")
		if len(host) <=0 {
			host = "localhost:6379"
		}
		pool := os.Getenv("REDIS_POOL")
		poolInt, err := strconv.Atoi(pool)
		if len(pool) <=0 || err != nil || poolInt <= 0{
			poolInt = 20
		}
		c.cache = cache.New(redis.NewClient(&redis.Options{
			Addr:     host,
			PoolSize: poolInt,
		}))
	}
	return c.cache
}