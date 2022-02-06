package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe("159.223.1.135:80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello team!")
}
