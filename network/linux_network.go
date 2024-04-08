package network

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const sysnet_path = "/sys/class/net/"

const (
	rx_bytes_path   = sysnet_path + "%s/statistics/rx_bytes"
	rx_packets_path = sysnet_path + "%s/statistics/rx_packets"
	rx_dropped_path = sysnet_path + "%s/statistics/rx_dropped"
	rx_errors_path  = sysnet_path + "%s/statistics/rx_errors"

	tx_bytes_path   = sysnet_path + "%s/statistics/tx_bytes"
	tx_packets_path = sysnet_path + "%s/statistics/tx_packets"
	tx_dropped_path = sysnet_path + "%s/statistics/tx_dropped"
	tx_errors_path  = sysnet_path + "%s/statistics/tx_errors"
)

// handleError handles all errors with the given message.
func handleError(err error, text string) {
	if err != nil {
		panic("\033[91m" + text + "\033[0m: " + err.Error())
	}
}

// readNetFiles reads data from the given file path and converts it to an integer.
func readNetFiles(fullPath string) (float64, error) {
	data, err := os.ReadFile(fullPath)
	if err != nil {
		handleError(err, "error in opening "+fullPath)
		return 0.0, err
	}
	dataString := string(data)
	data_replaced := strings.Replace(dataString, "\n", "", 1)
	result, err := strconv.ParseFloat(data_replaced, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// rx and tx FunctionBuilder returns a function that reads from the specified RX and TX file for the given interface.
func functionBuilder(pathFormat string) func(string) float64 {
	return func(interfaceName string) float64 {
		fullPath := fmt.Sprintf(pathFormat, interfaceName)
		result, err := readNetFiles(fullPath)
		handleError(err, "error reading "+fullPath)
		return result
	}
}

// Exports:

// listInterfaces returns a slice containing the names of all network interfaces.
func ListInterfaces() []string {
	files, err := os.ReadDir(sysnet_path)
	handleError(err, "error opening dir "+sysnet_path)

	var allInterfaces []string
	for _, file := range files {
		allInterfaces = append(allInterfaces, file.Name())
	}

	return allInterfaces
}

// Receive functions
var (
	RXBytes   = functionBuilder(rx_bytes_path)
	RXPackets = functionBuilder(rx_packets_path)
	RXDropped = functionBuilder(rx_dropped_path)
	RXErrors  = functionBuilder(rx_errors_path)
)

// Transmit functions
var (
	TXBytes   = functionBuilder(tx_bytes_path)
	TXPackets = functionBuilder(tx_packets_path)
	TXDropped = functionBuilder(tx_dropped_path)
	TXErrors  = functionBuilder(tx_errors_path)
)
