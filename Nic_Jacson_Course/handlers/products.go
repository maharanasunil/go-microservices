package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/sunilmaharana/data"
)

type Products struct {
	l *log.Logger
}

// NewProducts creates a product handler with given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// 1) getProducts returns the products from the data source
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	// d, err := json.Marshal(lp) // Easy Way

	err := lp.ToJSON(rw) // Using the encoder approach serialize the list to json

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	// rw.Write(d)
}

// 2) addProducts adds a product to our list of products
func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

// 3) updateProducts updates the product details for a given id
func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		// return
	}
}

// ServeHTTP is the main entry pointfor the handler and satisfies the http.Handler
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// 1) GET Method - to get the data
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// 2) POST Method - to add the data
	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	// 3) PUT Method - to update the data
	if r.Method == http.MethodPut {
		p.l.Println("PUT called", r.URL.Path)
		// Expect the id in the uri
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		p.l.Println("regexp.MustCompile() returns: ", r)
		// &{PUT /2 HTTP/1.1 1 1 map[Accept:[*/*]
		// Content-Length:[58]
		// Content-Type:[application/x-www-form-urlencoded]
		// User-Agent:[curl/7.78.0]] 0xc00005c200 <nil> 58 [] false
		// localhost:9090 map[] map[] <nil> map[] 127.0.0.1:56460 /2 <nil> <nil> <nil> / 0xc000130190 0xc00005e180 [] map[]}

		p.l.Println("FindAllStringSubmatch() returns: ", g) // [[/2 2]]
		if len(g) != 1 {
			p.l.Println("Invalid URL: More than one id")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URL: More than one capture group")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString) // convert the id to string type
		if err != nil {
			p.l.Println("Invalid URL: Unable to convert the number", idString)
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		p.l.Println("got id:", id)

		p.updateProducts(id, rw, r)
		return

	}

	// Catch all - if no method satisfies then return error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
