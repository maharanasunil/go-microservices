package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
	"github.com/sunilmaharana/handlers"
)

// env is a module created by nicholas to handle env variables
var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind Address for the server")

func main() {

	env.Parse()

	// Prefixes every log with product-api
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	// We use gorilla mux package now
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	// Register handler to my server via HandleFunc
	getRouter.HandleFunc("/", ph.GetProducts) // products handler (We capitalized the func name and it became public)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	// Use Middleware to validate JSON
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	// Use Middleware to validate JSON
	postRouter.Use(ph.MiddlewareValidateProduct)

	// It is just a type in server
	s := http.Server{
		Addr:         *bindAddress,
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server in a Goroutine
	// This piece of code starts your web server in the background, so that the main program can do something else (like listen for shutdown signals)
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe() // starts your HTTP server on the port you defined earlier
		if err != nil {
			l.Fatal(err) // Logs error and exits immediately.
		}

	}()

	// Creates a channel that receives OS signals.
	// chan = a channel — a Go feature for sending and receiving messages between goroutines
	// make() = built-in function to create slices, maps, or channels.
	sigChan := make(chan os.Signal)

	// Register the Channel for OS Signals
	signal.Notify(sigChan, os.Interrupt) // You hit Ctrl+C and signal is sent to sigChan
	signal.Notify(sigChan, os.Kill)      // Force Kill Signal

	sig := <-sigChan // <- is the receive operator for channels. which means: Wait until something is sent into sigChan
	l.Println("Received Terminate, graceful Shutdown!", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second) // Returns a context that cancels after a certain duration.
	// It takes a context as a parameter
	s.Shutdown(tc)

}
