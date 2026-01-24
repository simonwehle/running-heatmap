package cmd

import (
	"log"
	"net/http"
)

func Execute() {
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)

	addr := ":8000"
	log.Println("Serving on http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}