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
	message string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		fmt.Printf("Sleeping for %v seconds...\n", i)
		time.Sleep(time.Duration(i) * time.Second)
	}

	handler := &handler{
		message: message,
	}

	log.Println("Listening on", port)
	http.ListenAndServe(port, handler)
}
