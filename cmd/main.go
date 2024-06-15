package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/crud/configs/db"
	"example.com/crud/configs/env"
	"example.com/crud/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	envConfig := env.GetEnv()
	port := envConfig.AppPort
	print("port=============", port)
	// addrs := fmt.Sprintf(":%s", port)
	r := gin.Default()

	_, err := db.ConnectDatabase(envConfig)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	// Set routes
	routes.SetRoutes(r)

	// e := r.Run(":5000") // listen and serve on
	// if e != nil {
	// 	panic(e)
	// }
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	// Notify the channel when a SIGINT or SIGTERM is received
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Block until a signal is received
	<-quit

	// Create a context with a timeout for the graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Perform the graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown: %v\n", err)
	}
}
