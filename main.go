package main

import (
	"context"
	"fmt"

	"github.com/yash91989201/go_microservice/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
}
