package main

import (
	"fmt"

	"github.com/faysalahmed-dev/go-microservice-order-api/app"
)

func main() {
	a := app.App{}
	err := a.StartApp("localhost:3000");
	if err != nil {
		fmt.Println(err)
	}
}