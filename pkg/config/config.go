package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Config holds the application configuration, including logging settings
type Config struct {
	// Add other configuration fields as necessary
	LogLevel           string   `json:"logLevel"`           // LogLevel can be "debug", "info", "warn", "error"
	LogFormat          string   `json:"logFormat"`          // LogFormat can be "json" or "text"
	LogFile            string   `json:"logFile"`            // Optional: Path to log file, if logging to a file
	LogTimeFormat      string   `json:"logTimeFormat"`      // Optional: Custom time format for log messages
	WorkerPoolSize     int      `json:"workerPoolSize"`     // Number of workers in the worker pool
	HTTPTimeout        int      `json:"httpTimeout"`        // HTTP client timeout in seconds
	Proxies            []string `json:"proxies"`            // List of proxies
	DBConnectionString string   `json:"dbConnectionString"` // Database connection string
}

// LoadConfig loads the application configuration from a file, environment variables, or any other source
func LoadConfig() (*Config, error) {

	cfg := &Config{
		LogLevel:           "info",
		LogFormat:          "json",
		LogFile:            "", // Empty means logging to stdout
		LogTimeFormat:      "2006-01-02 03:04:05 PM",
		WorkerPoolSize:     5000, // Can be adjusted based on application requirements
		HTTPTimeout:        15,
		DBConnectionString: "your_db_connection_string_here",
	}

	// Load proxies from file
	proxies, err := LoadProxies("proxies.txt")
	if err != nil {
		return nil, err // Handle error appropriately
	}
	cfg.Proxies = proxies

	return cfg, nil
}

// LoadProxies reads the proxies from a newline-separated file.
func LoadProxies(filePath string) ([]string, error) {
	file, err := os.Open("proxies.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open proxies.txt: %w", err)
	}
	defer file.Close()

	var proxies []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ":")

		if len(fields) != 4 {
			continue
		}

		ip := fields[0]
		port := fields[1]
		username := fields[2]
		password := fields[3]

		proxy := "http://" + username + ":" + password + "@" + ip + ":" + port
		proxies = append(proxies, proxy)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error while reading proxies.txt: %w", err)
	}

	return proxies, nil
}
