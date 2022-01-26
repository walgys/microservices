package handlers

import (
	"log"
	"net/http"

	"github.com/walgys/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		p.getProducts(rw,r)
		return
	}

	if r.Method == http.MethodPut {
		p.updateProducts(rw,r)
		return
	}


	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request){
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err!= nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

func (p *Products) updateProducts(rw http.ResponseWriter, r *http.Request){
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err!= nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}