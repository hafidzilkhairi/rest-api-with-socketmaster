package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func fetchTheApi() {
	// set the time of log prefix to show the microsecond
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// defining the url and the method of api
	req, err := http.NewRequest(http.MethodGet, "http://localhost:10000/", nil)
	if err != nil {
		log.Println("[http.NewRequest]", err)
		return
	}

	// do send a request to the api server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[client.Do]", err)
		return
	}

	defer resp.Body.Close()

	// read response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[json.NewDecoder]", err)
		return
	}

	// print the response
	log.Println("Data: ", string(data))
}

func main() {
	// infinite loop, press ctrl-c to exit whenever you want
	for {
		fetchTheApi()
		time.Sleep(500 * time.Millisecond)
	}
}
