package main

import (
	"bufio" // reading & writing of file and input
	"fmt"
	"os"            // Operating-system functionality
	"path/filepath" // Working with file/directory paths
	"sort"          // sorting slices
	"strings"       // manipulating strings
)

// reating a new type called LogEntry - it represents one log entry.
// struct allows to group multiple related pieces of data together.
type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
	Source    string
}

func main() {

	// Read all files and folders inside the "logs" directory
	entries, err := os.ReadDir("Log-file-parser/logs")

	// Check if the directory was read successfully
	if err != nil {
		fmt.Println("Error reading log directory:", err)
		return
	}

	// var named logs that can hold multiple LogEntry val.
	var logs []LogEntry

	// Go through every entry inside the log directory
	// _ -index, ignored, entry - current directory entry,  entries - all directory entries
	for _, entry := range entries {

		// checking if the item is directory, if yes skip it
		if entry.IsDir() {
			continue
		}

		// Only process .log files
		if filepath.Ext(entry.Name()) != ".log" {
			continue
		}

		// Create the complete path to the log file
		filePath := filepath.Join("Log-file-parser/logs", entry.Name())

		// Read the log file
		fileLogs, err := readLogFile(filePath, entry.Name())

		// Check if the file was read successfully
		if err != nil {
			fmt.Println("Error reading", entry.Name(), ":", err)
			continue
		}

		// Add the logs from this file to our main logs slice
		// The ... means - Take all the elements inside fileLogs and pass them individually.
		logs = append(logs, fileLogs...)
	}

	// Sort all logs by timestamp
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].Timestamp < logs[j].Timestamp
	})

	// Print the aggregated logs
	fmt.Println("AGGREGATED LOGS")

	// looping through every log
	for _, log := range logs {
		fmt.Printf(
			"[%s] %s %-5s %s\n",
			log.Source,
			log.Timestamp,
			log.Level,
			log.Message,
		)
	}

	// Take all the logs in logs and write them into a file called aggregated.log.
	err = writeAggregatedLogs(logs, "aggregated.log")

	if err != nil {
		fmt.Println("Error writing aggregated log:", err)
		return
	}

	fmt.Println("\nAggregated logs saved to aggregated.log")
}

// readLogFile reads one log file and returns its log entries
func readLogFile(filePath string, source string) ([]LogEntry, error) {

	// Open the log file
	file, err := os.Open(filePath)

	// Check if the file opened successfully
	if err != nil {
		return nil, err
	}

	// Close the file when this function finishes
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Store the logs from this file
	var logs []LogEntry

	// Read every line
	for scanner.Scan() {

		// Get the current line
		line := scanner.Text()

		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Split the line into parts
		parts := strings.SplitN(line, " ", 4)

		// Make sure the line has the expected format
		if len(parts) < 4 {
			fmt.Println("Skipping invalid log line:", line)
			continue
		}

		// Create a LogEntry
		log := LogEntry{
			Timestamp: parts[0] + " " + parts[1],
			Level:     parts[2],
			Message:   parts[3],
			Source:    source,
		}

		// Add the log to the slice
		logs = append(logs, log)
	}

	// Check if scanner stopped because of an error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

// writeAggregatedLogs writes all logs to an output file
func writeAggregatedLogs(logs []LogEntry, filePath string) error {

	// Create the output file
	file, err := os.Create(filePath)

	// Check if the file was created successfully
	if err != nil {
		return err
	}

	// Close the file when the function finishes
	defer file.Close()

	// Write every log entry
	for _, log := range logs {

		fmt.Fprintf(
			file,
			"[%s] %s %-5s %s\n",
			log.Source,
			log.Timestamp,
			log.Level,
			log.Message,
		)
	}

	return nil
}
