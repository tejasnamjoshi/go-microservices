package main

import (
	"context"
	"go-microservices/product-images/files"
	"go-microservices/product-images/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api-gorilla", log.LstdFlags)
	basePath := "./imagestore"
	// create the handlers
	stor, err := files.NewLocal(basePath, 1024*1000*5)
	if err != nil {
		l.Panic(err)
	}

	fh := handlers.NewFiles(stor, l)

	// create a new server mux and register the handlers
	sm := mux.NewRouter()
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+.[a-z]{3}}", fh.ServeHTTP)
	ph.HandleFunc("/images", fh.SaveFileMultiPart)

	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z]+.[a-z]{3}}", http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))))

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	s := &http.Server{
		Addr:         ":9090",
		Handler:      ch(sm),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
