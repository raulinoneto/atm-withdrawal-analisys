package withdrawal

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type (
	ServiceResponse struct {
		Amount float64     `json:"amount"`
		Coins  map[int]int `json:"coins"`
	}

	// Cacher is a interface to provides to storage the values for best performance
	Cacher interface {
		//Get will return the stored coin count based in amount
		Get(ctx context.Context, amount int) map[int]int
		//Set will save the coin count based in amount
		Set(ctx context.Context, amount int, coinCount map[int]int)
	}
)

func (res *ServiceResponse) String() string {
	bytes, err := json.Marshal(res)
	if err != nil {
		logrus.Warn("trying to sent wrong response format "+ err.Error())
	}
	return string(bytes)
}