package app

import (
	"fmt"
	"net/http"

	route "github.com/faysalahmed-dev/go-microservice-order-api/routes"
)

type App struct {}

func (a *App) StartApp(host string) error {
	r:= route.LoadRoutes()
	err:= http.ListenAndServe(host, r)
	if err != nil {
		return fmt.Errorf("something went wrong, %v", err)
	}
	return nil
}