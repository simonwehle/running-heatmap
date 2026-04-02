package tiles

import (
	"net/http"
	"path/filepath"
	"strings"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, "./tiles/heatmap.mbtiles")
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("./tiles", r.URL.Path)
	
	cleanPath := filepath.Clean(filePath)
	if !strings.HasPrefix(cleanPath, filepath.Clean("./tiles")) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	
	if filepath.Ext(filePath) == ".mbtiles" {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	
	http.ServeFile(w, r, filePath)
}
