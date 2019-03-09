package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", res)

	err := server.ListenAndServe()

	if err != nil {
		log.Println(err)
	}
}

func res(w http.ResponseWriter, r *http.Request) {
	log.Println("receive request")
	fmt.Fprintf(w, "Hello Docker!!")
}
