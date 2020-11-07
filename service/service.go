package service

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// Service is the main logic of application
type Service struct {
	cache Cacher
}

// New Instantiate a Service pointer
func New(cache Cacher) *Service {
	return &Service{cache: cache}
}

// ProcessAmount give how many coins will be used at ATM Withdrawal
func (svc *Service) ProcessAmount(ctx context.Context, amount int) *ServiceResponse {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logger := logrus.WithContext(ctx).WithField("amount", amount).WithTime(time.Now().Add(time.Second))
	logger.Info("Initialize ProcessAmount")
	res := &ServiceResponse{
		Amount: float64(amount),
		Coins: make(map[int]int),
	}
	if coinCount := svc.cache.Get(ctx, amount); coinCount != nil {
		res.Coins = coinCount
		return res
	}
	for amount > 0 {
		amount = setCoin(logger, amount, res.Coins)
	}
	logger.WithField("coin_count", res.Coins).Info("coin Count Result")
	go svc.cache.Set(ctx, amount, res.Coins)
	return res
}

// setCoin Calculate the amount based in knew coins
func setCoin(logger *logrus.Entry, amount int, coinCount map[int]int) int {
	switch {
	case amount%Fifty == 0:
		amount -= Fifty
		coinCount[Fifty]++
		logger.Info(fmt.Sprintf("has %d fifty coins", coinCount[Fifty]))
	case amount%Ten == 0:
		amount -= Ten
		coinCount[Ten]++
		logger.Info(fmt.Sprintf("has %d ten coins", coinCount[Ten]))
	case amount%Five == 0:
		amount -= Five
		coinCount[Five]++
		logger.Info(fmt.Sprintf("has %d five coins", coinCount[Five]))
	case amount%One == 0:
		amount -= One
		coinCount[One]++
		logger.Info(fmt.Sprintf("has %d one coins", coinCount[One]))
	}
	return amount
}
