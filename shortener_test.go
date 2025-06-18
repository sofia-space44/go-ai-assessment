package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	InitializeStorage()
	
	r := gin.New()
	r.POST("/shorten", CreateShortURL)
	r.GET("/:code", RedirectToOriginal)
	r.GET("/analytics/:code", GetAnalytics)
	
	return r
}

func TestCreateShortURL_BasicFunctionality(t *testing.T) {
	router := setupTestRouter()
	
	requestBody := ShortenRequest{
		URL: "https://example.com",
	}
	jsonBody, _ := json.Marshal(requestBody)
	
	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusCreated, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	
	assert.Contains(t, response, "short_url")
	assert.Contains(t, response, "short_code")
	assert.Equal(t, "https://example.com", response["original_url"])
}

func TestCreateShortURL_InvalidURL(t *testing.T) {
	router := setupTestRouter()
	
	requestBody := ShortenRequest{
		URL: "invalid-url",
	}
	jsonBody, _ := json.Marshal(requestBody)
	
	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestReservedKeywords(t *testing.T) {
	testCases := []string{"admin", "api", "health", "ADMIN", "Api"}
	
	for _, keyword := range testCases {
		t.Run("keyword_"+keyword, func(t *testing.T) {
			assert.True(t, isReservedKeyword(keyword), "Expected %s to be reserved", keyword)
		})
	}
}

// TODO: Task 3 - Add comprehensive tests for:
// 1. Custom alias creation and validation
// 2. Reserved keyword blocking
// 3. Alias expiration logic
// 4. Analytics accuracy (including the bug fix)
// 5. Edge cases for alias conflicts

func TestCustomAliasCreation_TODO(t *testing.T) {
	// TODO: Test custom alias functionality
	// This test should cover:
	// - Valid custom alias creation
	// - Alias validation rules
	// - Reserved keyword rejection
	// - Duplicate alias handling
	
	t.Skip("TODO: Implement custom alias tests")
}

func TestAliasExpiration_TODO(t *testing.T) {
	// TODO: Test alias expiration
	// This test should cover:
	// - Aliases expire after 30 days
	// - Expired aliases return 404
	// - Cleanup of expired aliases
	
	t.Skip("TODO: Implement expiration tests")
}

func TestAnalyticsBugFix_TODO(t *testing.T) {
	// TODO: Test analytics accuracy
	// This test should verify:
	// - No double counting from same IP
	// - Correct unique vs total click counts
	// - Analytics work with custom aliases
	
	t.Skip("TODO: Implement analytics bug fix tests")
}
