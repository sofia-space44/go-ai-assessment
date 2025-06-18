package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type URLMapping struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	CustomAlias string    `json:"custom_alias,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	ClickCount  int       `json:"click_count"`
	IsActive    bool      `json:"is_active"`
}

type ShortenRequest struct {
	URL         string `json:"url" binding:"required"`
	CustomAlias string `json:"custom_alias,omitempty"`
}

var urlMappings = make(map[string]*URLMapping)
var aliasMappings = make(map[string]*URLMapping) // Maps custom aliases to URLMapping

const dataFile = "url_data.json"

// Reserved keywords that cannot be used as custom aliases
var reservedKeywords = []string{
	"admin", "api", "health", "analytics", "dashboard", 
	"login", "logout", "register", "settings", "help",
	"about", "contact", "privacy", "terms", "docs",
}

func InitializeStorage() {
	loadData()
}

func loadData() {
	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		// File doesn't exist, start with empty data
		return
	}
	
	var data struct {
		URLMappings   map[string]*URLMapping `json:"url_mappings"`
		AliasMappings map[string]*URLMapping `json:"alias_mappings"`
	}
	
	if err := json.Unmarshal(file, &data); err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		return
	}
	
	urlMappings = data.URLMappings
	aliasMappings = data.AliasMappings
	
	if urlMappings == nil {
		urlMappings = make(map[string]*URLMapping)
	}
	if aliasMappings == nil {
		aliasMappings = make(map[string]*URLMapping)
	}
}

func saveData() error {
	data := struct {
		URLMappings   map[string]*URLMapping `json:"url_mappings"`
		AliasMappings map[string]*URLMapping `json:"alias_mappings"`
	}{
		URLMappings:   urlMappings,
		AliasMappings: aliasMappings,
	}
	
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	
	return ioutil.WriteFile(dataFile, jsonData, 0644)
}

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 7
	
	result := make([]byte, length)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	
	return string(result)
}

func isReservedKeyword(alias string) bool {
	lowerAlias := strings.ToLower(alias)
	for _, keyword := range reservedKeywords {
		if lowerAlias == keyword {
			return true
		}
	}
	return false
}

func CreateShortURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	
	// Basic URL validation
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL must start with http:// or https://"})
		return
	}
	
	// TODO: Task 1 - Implement custom alias validation and creation
	// This is incomplete - candidates need to add:
	// 1. Check if custom alias is provided
	// 2. Validate custom alias (not reserved, not already taken, proper format)
	// 3. Handle expiration for custom aliases (30 days from creation)
	// 4. Create mapping in aliasMappings
	
	// For now, only generate random short codes
	shortCode := generateShortCode()
	
	// Ensure uniqueness
	for urlMappings[shortCode] != nil {
		shortCode = generateShortCode()
	}
	
	mapping := &URLMapping{
		ID:          fmt.Sprintf("url_%d", time.Now().UnixNano()),
		OriginalURL: req.URL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now(),
		ClickCount:  0,
		IsActive:    true,
	}
	
	urlMappings[shortCode] = mapping
	
	if err := saveData(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL mapping"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"short_url":  fmt.Sprintf("http://localhost:8080/%s", shortCode),
		"short_code": shortCode,
		"original_url": req.URL,
	})
}

func RedirectToOriginal(c *gin.Context) {
	code := c.Param("code")
	
	// First check regular short codes
	mapping := urlMappings[code]
	
	// TODO: Task 1 - Also check custom aliases

	
	if mapping == nil || !mapping.IsActive {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}
	
	// TODO: Task 1 - Check if custom alias has expired

	
	// Record analytics - THIS HAS A BUG (Task 2)
	RecordClick(code, c.ClientIP())
	
	c.Redirect(http.StatusMovedPermanently, mapping.OriginalURL)
}
