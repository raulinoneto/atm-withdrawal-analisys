package main

import (
	"github.com/raulinoneto/atm-withdrawal-analisys/config/container"
	"github.com/raulinoneto/atm-withdrawal-analisys/config/routes"
)

func main() {
	c := new(container.Container)
	server := c.GetServer(routes.GetRoutes(c), nil)
	server.Run()
}
