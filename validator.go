package main

import (
	"regexp"
	"strings"
	"time"
)

// TODO: Task 1 - Implement these validation functions for custom aliases

func isValidCustomAlias(alias string) bool {
	// TODO: Implement validation rules:
	// - Must be 3-20 characters
	// - Only alphanumeric and hyphens
	// - Cannot start or end with hyphen
	// - Cannot be a reserved keyword
	
	return false // Placeholder - candidates need to implement
}

func isAliasAvailable(alias string) bool {
	// TODO: Check if alias is already taken in aliasMappings
	return false // Placeholder - candidates need to implement
}

func createCustomAliasMapping(originalURL, alias string) *URLMapping {
	// TODO: Create URLMapping with custom alias
	// - Set ExpiresAt to 30 days from now
	// - Add to both urlMappings and aliasMappings
	
	return nil // Placeholder - candidates need to implement
}

func cleanupExpiredAliases() {
	// TODO: Task 1 - Implement cleanup of expired custom aliases
	// This should be called periodically or on each request
	
	now := time.Now()
	_ = now // Prevent unused variable error
	
	// Placeholder - candidates need to implement
}

// Helper function for alias validation
func validateAliasFormat(alias string) error {
	if len(alias) < 3 || len(alias) > 20 {
		return fmt.Errorf("alias must be between 3 and 20 characters")
	}
	
	// Only alphanumeric and hyphens
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]+$", alias)
	if !matched {
		return fmt.Errorf("alias can only contain letters, numbers, and hyphens")
	}
	
	// Cannot start or end with hyphen
	if strings.HasPrefix(alias, "-") || strings.HasSuffix(alias, "-") {
		return fmt.Errorf("alias cannot start or end with a hyphen")
	}
	
	return nil
}
