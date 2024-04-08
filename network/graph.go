package network

import (
	"bytes"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

type NetworkData struct {
	Timestamps     []time.Time
	DownloadSpeeds []float64
	UploadSpeeds   []float64
}

func (n *NetworkData) AddData(timestamp time.Time, downloadSpeed, uploadSpeed float64) {
	// Append new data point
	n.Timestamps = append(n.Timestamps, timestamp)
	n.DownloadSpeeds = append(n.DownloadSpeeds, downloadSpeed)
	n.UploadSpeeds = append(n.UploadSpeeds, uploadSpeed)

}

// RenderChartPNG returns the PNG chart as a byte slice.
func (n *NetworkData) RenderChartPNG() ([]byte, error) {
	// Create the download series.
	downloadSeries := chart.TimeSeries{
		Name: "Download",
		XValues: n.Timestamps,
		YValues: n.DownloadSpeeds,
		Style: chart.Style{
			StrokeColor: chart.ColorGreen,
			FillColor:   chart.ColorGreen.WithAlpha(64),
		},
	}

	// Create the upload series.
	uploadSeries := chart.TimeSeries{
		Name:    "Upload",
		XValues: n.Timestamps,
		YValues: n.UploadSpeeds,
		Style: chart.Style{
			StrokeColor: chart.ColorBlue,
			FillColor:   chart.ColorBlue.WithAlpha(64),
		},
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:           "Last 30 Minutes",
			ValueFormatter: func(v interface{}) string { return "-" }, // This will remove the x-axis labels
		},
		Background: chart.Style{
			FillColor: chart.ColorLightGray, // Light gray background
		},
		YAxis: chart.YAxis{
			Name: "Mbit/s",
		},
		Series: []chart.Series{
			downloadSeries,
			uploadSeries,
		},
	}

	var buffer bytes.Buffer
	err := graph.Render(chart.PNG, &buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
