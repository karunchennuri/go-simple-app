package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// IndexHandler returns a simple message
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I am Go App!")
}

func main() {
	http.HandleFunc("/", IndexHandler)

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "8888"
	}
	log.Println(http.ListenAndServe(":"+port, nil))
}
