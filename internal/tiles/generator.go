package tiles

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"running-heatmap/internal/parser"
)

type GeoJSONFeature struct {
	Type       string                 `json:"type"`
	Geometry   GeoJSONGeometry        `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type GeoJSONGeometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type FeatureCollection struct {
	Type     string          `json:"type"`
	Features []GeoJSONFeature `json:"features"`
}

type LineSegment struct {
	Coordinates [][]float64
	Count       int
}

var (
	cachedGeoJSON *FeatureCollection
	cacheMutex    sync.Mutex
)

func Generate() error {
	log.Println("Starting heatmap generation...")

	lineSegments, err := extractLineSegments()
	if err != nil {
		return err
	}
	log.Printf("Extracted %d line segments", len(lineSegments))

	geojson := segmentsToGeoJSON(lineSegments)
	log.Printf("Created GeoJSON with %d line features", len(geojson.Features))

	cacheMutex.Lock()
	cachedGeoJSON = &geojson
	cacheMutex.Unlock()

	log.Println("✓ Heatmap generation complete")
	return nil
}

func extractLineSegments() ([]LineSegment, error) {
	var allSegments []LineSegment

	files, err := os.ReadDir("./assets")
	if err != nil {
		return nil, err
	}

	processedFiles := 0
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".gpx") {
			continue
		}

		filePath := filepath.Join("./assets", file.Name())
		gpx, err := parser.ParseGPXFile(filePath)
		if err != nil {
			log.Printf("Warning: skipping %s: %v", file.Name(), err)
			continue
		}

		for _, track := range gpx.Tracks {
			for _, segment := range track.Segments {
				if len(segment.Points) < 2 {
					continue
				}

				coords := make([][]float64, len(segment.Points))
				for i, point := range segment.Points {
					coords[i] = []float64{point.Lon, point.Lat}
				}

				allSegments = append(allSegments, LineSegment{
					Coordinates: coords,
					Count:       1,
				})
			}
		}
		processedFiles++
	}

	log.Printf("Processed %d GPX files", processedFiles)
	return allSegments, nil
}

func segmentsToGeoJSON(segments []LineSegment) FeatureCollection {
	fc := FeatureCollection{
		Type:     "FeatureCollection",
		Features: []GeoJSONFeature{},
	}

	for i, segment := range segments {
		feature := GeoJSONFeature{
			Type: "Feature",
			Geometry: GeoJSONGeometry{
				Type:        "LineString",
				Coordinates: segment.Coordinates,
			},
			Properties: map[string]interface{}{
				"count": segment.Count,
				"id":    i,
			},
		}
		fc.Features = append(fc.Features, feature)
	}

	return fc
}

func GetHeatmapGeoJSON(w http.ResponseWriter, r *http.Request) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if cachedGeoJSON == nil {
		http.Error(w, "Heatmap not ready", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cachedGeoJSON)
}
