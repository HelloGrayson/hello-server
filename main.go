package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type handler struct {
	message       string
	statusCode    int
	responseDelay time.Duration
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.responseDelay > 0 {
		log.Printf("Sleeping for %i seconds...\n", h.responseDelay)
		time.Sleep(h.responseDelay * time.Second)
	}
	if h.statusCode != 200 {
		http.Error(w, http.StatusText(h.statusCode), h.statusCode)
		return
	}
	fmt.Fprint(w, h.message)
}

func main() {

	port := fmt.Sprintf(":%v", os.Getenv("HELLO_PORT"))
	if port == ":" {
		port = ":8080"
	}

	message := os.Getenv("HELLO_MESSAGE")
	if message == "" {
		message = "Hello World!"
	}

	sleep := os.Getenv("HELLO_SLEEP")
	if sleep != "" {
		i, err := strconv.Atoi(sleep)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Sleeping for %v seconds...\n", i)
		time.Sleep(time.Duration(i) * time.Second)
	}

	code := 200
	statusCodeStr := os.Getenv("HELLO_STATUS_CODE")
	if statusCodeStr != "" {
		statusCode, err := strconv.Atoi(statusCodeStr)
		if err != nil {
			log.Fatalf("HELLO_STATUS_CODE must be an int, err: %v", err)
		}
		if http.StatusText(statusCode) == "" {
			log.Fatalf("HELLO_STATUS_CODE must be a valid status code, got: %v", statusCode)
		}
		code = statusCode
	}

	responseDelay := 0
	responseDelayStr := os.Getenv("HELLO_RESPONSE_DELAY")
	if responseDelayStr != "" {
		var err error
		responseDelay, err = strconv.Atoi(responseDelayStr)
		if err != nil {
			log.Fatalf("HELLO_RESPONSE_DELAY must be an int, err: %v", err)
		}
	}

	handler := &handler{
		message:       message,
		statusCode:    code,
		responseDelay: time.Duration(responseDelay),
	}

	log.Println("Listening on", port)
	log.Println("Will respond with body", message, "and status code", code)
	http.ListenAndServe(port, handler)
}
