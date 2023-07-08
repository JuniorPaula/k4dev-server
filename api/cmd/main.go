package main

import (
	"fmt"
	"knowledge-api/internal/config"
	"knowledge-api/internal/router"
	"log"
	"net/http"
)

func main() {
	// config.InitEnv()
	config.InitEnv()

	fmt.Println("[::] Server running on port:", config.Port)
	r := router.HanlderRoutes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
