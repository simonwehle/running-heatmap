package cmd

import (
	"log"
	"net/http"
	"text/template"

	"running-heatmap/internal/files"
	"running-heatmap/internal/style"
	"running-heatmap/internal/tiles"
)

func Execute() {
	mapStyle := style.GetMapStyle()

	if err := tiles.Generate(); err != nil {
		log.Fatalf("Failed to generate tiles: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		tmpl, _ := template.ParseFiles("./web/index.html")
		tmpl.Execute(w, map[string]string{"MapStyle": mapStyle})
	})

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.Handle("/tiles/", http.StripPrefix("/tiles/", http.FileServer(http.Dir("./tiles"))))
	http.HandleFunc("/api/gpx", files.ListGPXFiles)
	http.HandleFunc("/api/heatmap", tiles.GetHeatmapGeoJSON)
	addr := ":8080"
	log.Println("Server started at port" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}