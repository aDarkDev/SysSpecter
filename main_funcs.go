package main

import (
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/valyala/fasthttp"
	
	"encoding/json"
	"syscall"
	"time"
	"fmt"
	
)

// Json converts any data or []string to JSON format
func Json(data interface{}) (string, error) {
	// Marshal the data to JSON with indentation for better readability
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func GetDiskUsage() (map[string]float64, error) {
	usageMap := make(map[string]float64)

	// Get a list of disk partitions
	partitions, err := disk.Partitions(true) // true to include all devices
	if err != nil {
		usageMap["error"] = -1
		return usageMap, err // handle error
	}

	// Iterate over partitions to get disk usage
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			// Skip or handle error
			continue
		}

		// Store usage percentage in the map with mountpoint as the key
		usageMap[partition.Mountpoint] = usage.UsedPercent
	}

	return usageMap, nil
}

func formatUptime(uptimeSecs int64) string {
	uptime := time.Duration(uptimeSecs) * time.Second
	days := uptime / (24 * time.Hour)
	uptime -= days * 24 * time.Hour
	hours := uptime / time.Hour
	uptime -= hours * time.Hour
	minutes := uptime / time.Minute

	return fmt.Sprintf("%d day(s) %d hour(s) %d min(s)", days, hours, minutes)
}


func GetUptime() (string,int64) {
	var sysinfo syscall.Sysinfo_t
	err := syscall.Sysinfo(&sysinfo)
	if err != nil {
		fmt.Println("Error:", err)
		return "",0
	}
	return formatUptime(sysinfo.Uptime),sysinfo.Uptime
}

func ctx_error_handler(ctx *fasthttp.RequestCtx, err error) {
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		fmt.Fprintf(ctx, "An error occurred: Check terminal errors.")
	}
}

func Contains(slice []string, toCheck string) bool {
	found := false
	for _, value := range slice {
		if value == toCheck {
			found = true
			break
		}
	}

	return found
}
