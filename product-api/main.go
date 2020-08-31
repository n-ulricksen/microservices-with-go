package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"scratch/microservices-with-go/product-api/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create handlers
	productsHandler := handlers.NewProducts(logger)

	// create mux, register handlers
	muxRouter := mux.NewRouter()

	getRouter := muxRouter.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productsHandler.GetProducts)

	putRouter := muxRouter.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProduct)
	putRouter.Use(productsHandler.MiddlewareProductValidation)

	postRouter := muxRouter.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productsHandler.AddProduct)
	postRouter.Use(productsHandler.MiddlewareProductValidation)

	// create a new server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      muxRouter,
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
