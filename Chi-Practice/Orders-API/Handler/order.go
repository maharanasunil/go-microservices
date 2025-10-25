package handler

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order")
}

// List all the elements that match a filter
func (o *Order) List(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("List all orders")
}

// Return an order by its id
func (o *Order) GetByID(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order by id")
}

// Update an order by its id
func (o *Order) UpdateByID(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by id")
}

// Delete an order by its id
func (o *Order) DeleteByID(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by id")
}
