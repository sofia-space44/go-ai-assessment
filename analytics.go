package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ClickRecord struct {
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	ShortCode string    `json:"short_code"`
}

type AnalyticsResponse struct {
	ShortCode    string        `json:"short_code"`
	OriginalURL  string        `json:"original_url"`
	TotalClicks  int           `json:"total_clicks"`
	UniqueClicks int           `json:"unique_clicks"`
	RecentClicks []ClickRecord `json:"recent_clicks"`
	CreatedAt    time.Time     `json:"created_at"`
}

var clickRecords = make([]ClickRecord, 0)
const analyticsFile = "analytics_data.json"

func init() {
	loadAnalytics()
}

func loadAnalytics() {
	file, err := ioutil.ReadFile(analyticsFile)
	if err != nil {
		return // File doesn't exist, start with empty data
	}
	
	if err := json.Unmarshal(file, &clickRecords); err != nil {
		fmt.Printf("Error loading analytics: %v\n", err)
	}
}

func saveAnalytics() error {
	jsonData, err := json.MarshalIndent(clickRecords, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(analyticsFile, jsonData, 0644)
}

// BUG: This function has a bug - it doesn't properly deduplicate IPs
// Task 2: Fix the double-counting issue
func RecordClick(shortCode string, clientIP string) {
	// INTENTIONAL BUG: This records every click without checking for duplicates from same IP
	// The bug is that we should only count unique IPs per day, not every single click
	record := ClickRecord{
		Timestamp: time.Now(),
		IP:        clientIP,
		ShortCode: shortCode,
	}
	
	clickRecords = append(clickRecords, record)
	
	// Update the mapping's click count
	if mapping := urlMappings[shortCode]; mapping != nil {
		mapping.ClickCount++
	}
	// TODO: Task 1 - Also check custom aliases
	
	saveAnalytics()
	saveData()
}

func GetAnalytics(c *gin.Context) {
	code := c.Param("code")
	
	mapping := urlMappings[code]
	// TODO: Task 1 - Also check custom aliases

	
	if mapping == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}
	
	// Get click records for this short code
	var urlClicks []ClickRecord
	uniqueIPs := make(map[string]bool)
	
	for _, record := range clickRecords {
		if record.ShortCode == code {
			urlClicks = append(urlClicks, record)
			uniqueIPs[record.IP] = true
		}
	}
	
	// Get recent clicks (last 10)
	recentClicks := urlClicks
	if len(recentClicks) > 10 {
		recentClicks = recentClicks[len(recentClicks)-10:]
	}
	
	response := AnalyticsResponse{
		ShortCode:    code,
		OriginalURL:  mapping.OriginalURL,
		TotalClicks:  mapping.ClickCount,
		UniqueClicks: len(uniqueIPs),
		RecentClicks: recentClicks,
		CreatedAt:    mapping.CreatedAt,
	}
	
	c.JSON(http.StatusOK, response)
}
