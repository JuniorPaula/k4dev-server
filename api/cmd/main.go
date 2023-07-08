package main

import (
	"fmt"
	"knowledge-api/internal/router"
	"log"
	"net/http"
)

func main() {
	//
	fmt.Println("[::] Server running on port: 8080")
	r := router.HanlderRoutes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), r))
}
