package files

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ListGPXFiles(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./assets")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var gpxFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".gpx") {
			gpxFiles = append(gpxFiles, filepath.Join("/assets", file.Name()))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gpxFiles)
}