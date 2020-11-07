package main

import (
	"fmt"
	"github.com/raulinoneto/atm-withdrawal-analisys/internal/httpserver"
	"net/http"
)

func main() {
	server := httpserver.New(&httpserver.Options{
		Middlewares: nil,
		Routes: []httpserver.Route{
			{
				Path:   "/test",
				Method: http.MethodGet,
				Handler: func(writer http.ResponseWriter, request *http.Request) error {
					//fmt.Fprint(writer, "test")
					fmt.Fprint(writer, request.FormValue("id"))
					return nil
				},
				Middlewares: nil,
			},
		},
		Port: "8080",
		Host: "",
	})
	server.Run()
}
