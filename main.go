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
	message := os.Getenv("HELLO_MESSAGE")

	handler := &handler{
		message: message,
	}

	log.Println("Listening on", port)
	http.ListenAndServe(port, handler)
}
