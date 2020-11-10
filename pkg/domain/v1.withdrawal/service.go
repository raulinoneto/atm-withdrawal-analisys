package withdrawal

import (
	"context"
	"fmt"
	"github.com/raulinoneto/atm-withdrawal-analisys/tools/logger"
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
	log := logger.New(ctx)
	log.WithField("amount", amount)
	log.Info("Initialize ProcessAmount")
	res := &ServiceResponse{
		Amount: float64(amount),
		Coins:  make(map[int]int),
	}
	if coinCount := svc.cache.Get(ctx, amount); coinCount != nil {
		res.Coins = coinCount
		return res
	}
	for amount > 0 {
		amount = setCoin(log, amount, res.Coins)
	}
	log.WithField("coin_count", res.Coins)
	log.Info("coin Count Result")
	svc.cache.Set(ctx, int(res.Amount), res.Coins)
	return res
}

// setCoin Calculate the amount based in knew coins
func setCoin(logger *logger.Logger, amount int, coinCount map[int]int) int {
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
