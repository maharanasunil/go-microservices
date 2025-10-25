package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", basicHandler)

	// Instantiate a http server
	server := &http.Server{
		Addr:    ":3000", // Port
		Handler: router,  // in chi the router type itself confirms to the http handler
		// http.HandlerFunc(basicHandler), // this function will be called when our server receives a request
	}

	// Call our server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen and serve", err)
		return
	}
}

func basicHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello, Sunil!"))
}
