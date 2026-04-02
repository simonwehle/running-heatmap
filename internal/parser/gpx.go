package parser

import (
	"encoding/xml"
	"os"
)

type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Tracks  []Track  `xml:"trk"`
}

type Track struct {
	Segments []TrackSegment `xml:"trkseg"`
}

type TrackSegment struct {
	Points []TrackPoint `xml:"trkpt"`
}

type TrackPoint struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
}

func ParseGPXFile(filepath string) (*GPX, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var gpx GPX
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&gpx); err != nil {
		return nil, err
	}

	return &gpx, nil
}

func (g *GPX) GetAllPoints() []TrackPoint {
	var points []TrackPoint
	for _, track := range g.Tracks {
		for _, segment := range track.Segments {
			points = append(points, segment.Points...)
		}
	}
	return points
}
