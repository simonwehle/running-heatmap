package cmd

import (
	"log"
	"net/http"

	"running-heatmap/internal/files"
)

func Execute() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/api/gpx", files.ListGPXFiles)
	addr := ":8080"
	log.Println("Server started at port" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}