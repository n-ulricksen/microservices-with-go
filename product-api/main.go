package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"scratch/microservices-with-go/product-api/handlers"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create handlers
	productsHandler := handlers.NewProducts(logger)

	// create mux, register handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/", productsHandler)

	// create a new server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server
	go func() {
		logger.Println("Server started on port 8080")
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// graceful shutdown
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)

	sig := <-signalChan
	logger.Println("Received terminate, graceful shutdown", sig)

	// Allow 30 seconds for graceful shutdown; forcefully close any handlers
	// still running after timeout duration.
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(timeoutContext)
}
