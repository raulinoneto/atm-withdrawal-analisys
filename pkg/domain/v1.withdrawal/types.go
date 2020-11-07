package v1_withdrawal

import "context"

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
