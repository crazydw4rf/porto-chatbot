package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crazydw4rf/porto-chatbot/api"
	"github.com/crazydw4rf/porto-chatbot/internal/config"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	host := cfg.APP_HOST
	if host == "" {
		host = "127.0.0.1"
	}

	port := cfg.APP_PORT
	if port == 0 {
		port = 3000
	}

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Server starting on http://%s:%d", host, port)

	if err := http.ListenAndServe(addr, http.HandlerFunc(api.Handler)); err != nil {
		log.Fatal(err)
	}
}
