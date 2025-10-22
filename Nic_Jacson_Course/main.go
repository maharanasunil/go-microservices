package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sunilmaharana/handlers"
)

func main() {

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Println("Hello World")
			d, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Ooops", http.StatusBadRequest)
				// w.WriteHeader(http.StatusBadRequest)
				// w.Write([]byte("Ooops"))
				return
			}
			// log.Printf("Data %s\n", d)
			fmt.Fprintf(w, "Hello %s", d)
		})
		http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
			log.Println("GoodBye World")
		})
		http.ListenAndServe(":9090", sm)
	*/

	// Prefixes every log with product-api
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Each handler gets the logger so they can log messages internally.
	// hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)
	ph := handlers.NewProducts(l)

	// '''ServeMux routes requests based on path'''

	sm := http.NewServeMux()
	// Register each handler to my server via servemux

	// sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/", ph) // products handler

	// It is just a type in server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Start Server in a Goroutine
	// This piece of code starts your web server in the background, so that the main program can do something else (like listen for shutdown signals)
	go func() {
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
