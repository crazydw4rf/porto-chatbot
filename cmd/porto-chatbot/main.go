package main

import (
	"fmt"

	"github.com/crazydw4rf/porto-chatbot/api"
)

func main() {
	api.LoadEnvVars()

	app := api.NewFiberApp()

	err := app.Listen(":3000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
