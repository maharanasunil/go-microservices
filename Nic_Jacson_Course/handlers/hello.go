package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// This is a struct to hold the pointer for logger
type Hello struct {
	l *log.Logger
}

// New function to return the logging data
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}

	// log.Printf("Data %s\n", d)
	fmt.Fprintf(rw, "Hello %s", d)
}
