package main

import (
	"context"
	"fmt"
	"knowledge-api/internal/config"
	"knowledge-api/internal/router"
	"knowledge-api/internal/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/cors"
)

func main() {
	// config.InitEnv()
	config.InitEnv()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// create the context to control of the exit
	_, cancel := context.WithCancel(context.Background())

	// init schedule stats
	go func() {
		schedule := utils.NewSchedule()
		schedule.StatsSchedule()
	}()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{config.FrontendURL},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: c.Handler(router.HanlderRoutes()),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error while start the server: %v", err)
		}
	}()

	// print out the port
	fmt.Println("[+] server running on port:", config.Port)

	// wait for the interrupt signal to gracefully shutdown the server with
	<-interrupt
	fmt.Println("server shutting down...")
	cancel()

	// a timeout of 5 seconds
	timeout := 5 * time.Second
	ctxShutDown, cancelShutdown := context.WithTimeout(context.Background(), timeout)
	defer func() {
		cancelShutdown()
	}()

	if err := server.Shutdown(ctxShutDown); err != nil {
		fmt.Println("error while shutingdown....", err)
		os.Exit(1)
	}

	fmt.Println("server exiting")
}
