package app

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v5"
)

type DashboardData struct {
	CronLogs        []LogEntry `json:"cron_logs"`
	TransferLogs    []LogEntry `json:"transfer_logs"`
	CreatedInvoices []LogEntry `json:"created_invoices"`
}

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

func DashboardHandler(c *echo.Context) error {
	data := syncLogsToJSON()
	saveToJSONDB(data)

	return c.File("view/dashboard.html")
}

func DashboardDataHandler(c *echo.Context) error {
	data := syncLogsToJSON()
	return c.JSON(http.StatusOK, data)
}

func syncLogsToJSON() DashboardData {
	return DashboardData{
		CronLogs:        parseLogFile("logs/cron_times.txt"),
		TransferLogs:    parseLogFile("logs/transfer.txt"),
		CreatedInvoices: parseLogFile("logs/created.txt"),
	}
}

func parseLogFile(filename string) []LogEntry {
	entries := []LogEntry{}
	file, err := os.Open(filename)
	if err != nil {
		return entries
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 3)
		if len(parts) >= 3 {
			entries = append(entries, LogEntry{
				Timestamp: parts[0] + " " + parts[1],
				Message:   parts[2],
			})
		}
	}
	return entries
}

func saveToJSONDB(data DashboardData) {
	file, _ := json.MarshalIndent(data, "", "  ")
	_ = os.MkdirAll("logs", 0755)
	_ = os.WriteFile("logs/db.json", file, 0644)
}
