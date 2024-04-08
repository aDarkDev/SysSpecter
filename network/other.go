package network

import (
	"bytes"
	"math"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func round(x float64, places int) float64 {
	adjustment := math.Pow(10, float64(places))
	// Adjust the number to the desired number of decimal places
	rounded := math.Round(x*adjustment) / adjustment
	return rounded
}

func shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(DefaultShell, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// Exports:

var DefaultPerSecondSize = float64(1024 * 1024)
var DefaultInterface = ""
var DefaultShell = "bash"
var IncomingPerSecond = 0.0
var OutgoingPerSecond = 0.0
var NetworkData_start = NetworkData{}
var StartedGraphTime = time.Now().Unix()

func CalculatePerSecond() {
	LastIncome := RXBytes(DefaultInterface)
	LastOutgo := TXBytes(DefaultInterface)

	for {
		IncomeBytes := RXBytes(DefaultInterface)
		OutGoBytes := TXBytes(DefaultInterface)
		IncomingPerSecond = round((IncomeBytes-LastIncome)/DefaultPerSecondSize, 3) * 8 // update v1.2 : changed to Mbit per second.
		OutgoingPerSecond = round((OutGoBytes-LastOutgo)/DefaultPerSecondSize, 3) * 8   // update v1.2 : changed to Mbit per second.
		LastIncome = IncomeBytes
		LastOutgo = OutGoBytes

		if time.Now().Unix()-StartedGraphTime > 30*60*60 /* reset graph after 30 minutes*/ { 
			StartedGraphTime = time.Now().Unix()
			NetworkData_start = NetworkData{}
		}

		NetworkData_start.AddData(time.Now(), IncomingPerSecond, OutgoingPerSecond)
		time.Sleep(time.Millisecond * 1050)
	}
}

func GetEstablishedConnections() []map[string]string {
	stdout, _, err := shellout("netstat -nt | grep ESTABLISHED")
	if err != nil {
		handleError(err, "please install net-tools first. sudo apt install net-tools")
	}
	ipPortRegex := regexp.MustCompile(`\b((?:\d{1,3}\.){3}\d{1,5})\b:(\d+)\b`)
	// Find all matches of IP addresses and ports in each line
	lines := strings.Split(stdout, "\n")
	var result []map[string]string
	for _, line := range lines {
		submatches := ipPortRegex.FindAllStringSubmatch(line, 1000)
		if len(submatches) == 2 {
			ip1 := submatches[0][1]
			port1 := submatches[0][2]
			ip2 := submatches[1][1]
			port2 := submatches[1][2]
			protocol := strings.Split(line, " ")[0] // update v1.2: add protocol of source and dest
			result = append(result, map[string]string{
				"source_ip":   ip1,
				"source_port": port1,
				"dest_ip":     ip2,
				"dest_port":   port2,
				"protocol":    protocol,
			})
		}
	}

	return result
}
