package routes

import (
	"github.com/raulinoneto/atm-withdrawal-analisys/config/container"
	"github.com/raulinoneto/atm-withdrawal-analisys/internal/httpserver"
	"net/http"
)

func GetRoutes(c *container.Container) []httpserver.Route {
	return []httpserver.Route{
		{
			Path:        "/v1/withdrawal",
			Method:      http.MethodGet,
			Handler:     c.GetHttpAdapter().WithdrawalHandler,
			Middlewares: nil,
		},
	}
}