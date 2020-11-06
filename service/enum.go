package service

type (
	Bill float64
	BillCount map[Bill]int
)

const (
	Fifty Bill = 50.00
	Ten   Bill = 10.00
	Five  Bill = 5.00
	One   Bill = 1.00
)
