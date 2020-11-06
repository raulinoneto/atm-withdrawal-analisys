package service

import "context"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (svc *Service) ProcessAmount(ctx context.Context, amount float64) BillCount{
	return nil
}

