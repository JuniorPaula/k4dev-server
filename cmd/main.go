package main

import (
	"fmt"
	"knowledge-api/internal/config"
	"knowledge-api/internal/router"
	"knowledge-api/internal/utils"
	"log"
	"net/http"
)

func main() {
	// config.InitEnv()
	config.InitEnv()

	fmt.Println("[::] Server running on port:", config.Port)
	r := router.HanlderRoutes()

	// init schedule stats
	go func() {
		schedule := utils.NewSchedule()
		schedule.StatsSchedule()
	}()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
