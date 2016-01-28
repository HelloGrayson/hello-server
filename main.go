package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	handler := &handler{
		message: message,
	}

	log.Println("Listening on", port)
	http.ListenAndServe(port, handler)
}
