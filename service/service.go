package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (svc *Service) ProcessAmount(ctx context.Context, amount int) map[int]int {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logger := logrus.WithContext(ctx).WithField("amount", amount).WithTime(time.Now().Add(time.Second))
	logger.Info("Initialize ProcessAmount")
	billCount := make(map[int]int)
	for amount > 0 {
		switch {
		case amount%Fifty == 0:
			amount -= Fifty
			billCount[Fifty]++
			logger.Info(fmt.Sprintf("has %d fifty bills", billCount[Fifty]))
		case amount%Ten == 0:
			amount -= Ten
			billCount[Ten]++
			logger.Info(fmt.Sprintf("has %d ten bills", billCount[Ten]))
		case amount%Five == 0:
			amount -= Five
			billCount[Five]++
			logger.Info(fmt.Sprintf("has %d five bills", billCount[Five]))
		case amount%One == 0:
			amount -= One
			billCount[One]++
			logger.Info(fmt.Sprintf("has %d one bills", billCount[One]))
		}
	}
	logger.WithField("bill_count",billCount).Info("Bill Count Result")
	return billCount
}
