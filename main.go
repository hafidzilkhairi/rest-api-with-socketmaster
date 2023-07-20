package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	// open a file with fd 3
	f := os.NewFile(uintptr(3), "listener")

	// set the file as listener
	listener, err := net.FileListener(f)
	if err != nil {
		log.Fatalln("[net.FileListener]", err)
	}

	// new server
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)

	// set server as the handler
	srv := http.Server{
		Handler: mux,
	}

	// run server at the background
	go func() {
		log.Println(srv.Serve(listener))
	}()

	// create a variable for system signal
	term := make(chan os.Signal)
	signal.Notify(term, syscall.SIGTERM)

	// <-term will wait for the signal
	log.Println("got signal", <-term)
	log.Println("shutting down...")

	// shutdown gracefully so all active connections won't close until done
	srv.Shutdown(context.Background())
}
