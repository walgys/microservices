package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/walgys/microservices/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/products", ph)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		ErrorLog: l,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func(){

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting Server %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 *time.Second)
	s.Shutdown(tc)
	
}