# Go AI-Assisted Assessment Challenge - URL Shortener

**Time Limit: 75 minutes**
**AI Tools: Encouraged and Required**

## Overview

You will extend an existing URL shortener service by implementing custom alias functionality, fixing a critical analytics bug, and writing comprehensive tests. This challenge tests your ability to work with AI tools effectively while maintaining code quality and following existing patterns.

## Setup Instructions

### Prerequisites
- Go 1.21 or higher
- Your preferred AI coding assistant (Cursor, GitHub Copilot, ChatGPT, Claude, etc.)

### Getting Started
1. Clone this repository
2. Navigate to the project directory
3. Install dependencies: `go mod download`
4. Run the server: `go run .`
5. Test the API: `curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url":"https://google.com"}'`

The server will start on `http://localhost:8080`

## Your Tasks (75 minutes total)

### Task 1: Custom Alias Feature Implementation (25 minutes)
**Implement the ability for users to specify custom aliases for their short URLs.**

**Requirements:**
- Allow users to provide a `custom_alias` in the shorten request
- Validate custom aliases using these rules:
  - Must be 3-20 characters long
  - Only alphanumeric characters and hyphens allowed
  - Cannot start or end with a hyphen
  - Cannot be a reserved keyword (admin, api, health, etc.)
  - Must be unique (not already taken)
- Custom aliases expire after 30 days of no usage
- Expired aliases should return 404 and be cleaned up
- Both regular short codes and custom aliases should work in redirects and analytics

**Files to modify:**
- `shortener.go` - Complete the TODO sections in `CreateShortURL` and `RedirectToOriginal`
- `validator.go` - Implement all the placeholder functions

**Acceptance Criteria:**
- POST `/shorten` accepts `custom_alias` parameter
- GET `/:code` works with both generated codes and custom aliases
- Reserved keywords are rejected with appropriate error messages
- Duplicate aliases are rejected
- Expired aliases return 404

### Task 2: Fix Analytics Double-Counting Bug (25 minutes)
**There is a critical bug in the analytics system that double-counts clicks from the same IP address.**

**The Problem:**
The `RecordClick` function in `analytics.go` records every single click, even multiple clicks from the same IP address. This inflates the click counts and makes analytics unreliable.

**Requirements:**
- Fix the analytics to only count unique IP addresses per day
- Same IP clicking multiple times in one day = 1 unique click for that day
- Same IP clicking on different days = separate unique clicks
- Maintain total click count (all clicks) and unique click count separately
- Ensure the fix works with both regular short codes and custom aliases

**Files to modify:**
- `analytics.go` - Fix the `RecordClick` function and update related logic

**Acceptance Criteria:**
- Multiple clicks from same IP on same day count as 1 unique click
- Same IP on different days counts as separate unique clicks
- Total clicks still show all clicks (for rate limiting, etc.)
- Analytics endpoint returns both `total_clicks` and `unique_clicks`

### Task 3: Comprehensive Testing (25 minutes)
**Write thorough tests for the new functionality and bug fixes.**

**Requirements:**
Write tests that cover:
- Custom alias creation with valid inputs
- Custom alias validation (all the edge cases)
- Reserved keyword rejection
- Duplicate alias handling
- Alias expiration logic
- Analytics accuracy (verifying the bug fix)
- IP deduplication in analytics

**Files to modify:**
- `shortener_test.go` - Replace TODO sections with comprehensive tests

**Acceptance Criteria:**
- All tests pass with `go test -v`
- Test coverage includes both success and failure cases
- Tests verify the analytics bug fix works correctly
- Tests check custom alias expiration behavior

## AI Usage Documentation Requirements

**CRITICAL: You must document your AI usage to receive full credit.**

### How to Document Your AI Usage:

#### Option 1: Conversation-Based AI (ChatGPT, Claude, etc.)
- Share your conversation link(s) or export/screenshot your full conversations
- Include the prompts you used and how you refined them
- Show how you iterated based on AI responses

#### Option 2: Inline AI Tools (Cursor, Copilot, etc.)
- Take screenshots of key AI interactions
- Document the prompts/comments you used to trigger AI suggestions
- Briefly describe how you modified AI-generated code

#### Option 3: Mixed Usage
- Combine both approaches above
- Clearly separate which AI tool was used for which parts

### Minimum Requirements:
- We need to see your actual prompts/interactions with AI
- Show at least 3-5 meaningful AI interactions
- Demonstrate how you guided the AI to produce better results
- Include any cases where you had to correct or modify AI output

### Submission Format:
Create a file called `AI_USAGE.md` with:
- Links to your AI conversations (if applicable)
- Screenshots of key interactions (if applicable)  
- Brief summary of how AI helped you complete each task
- Any challenges you faced with AI and how you overcame them

## API Endpoints

### POST /shorten
Create a new short URL with optional custom alias.

**Request:**
```json
{
  "url": "https://example.com",
  "custom_alias": "my-link"  // optional
}
```

**Response:**
```json
{
  "short_url": "http://localhost:8080/my-link",
  "short_code": "my-link",
  "original_url": "https://example.com"
}
```

### GET /:code
Redirect to the original URL (works with both generated codes and custom aliases).

### GET /analytics/:code  
Get analytics for a short URL.

**Response:**
```json
{
  "short_code": "abc123",
  "original_url": "https://example.com",
  "total_clicks": 10,
  "unique_clicks": 7,
  "recent_clicks": [...],
  "created_at": "2023-01-01T00:00:00Z"
}
```

## Code Structure

- `main.go` - API server setup and routing
- `shortener.go` - Core URL shortening logic  
- `analytics.go` - Click tracking and analytics (contains the bug to fix)
- `validator.go` - Validation helpers for custom aliases (mostly incomplete)
- `shortener_test.go` - Test suite (needs completion)

## Evaluation Criteria

You will be evaluated on:

1. **Task Completion (1-5)**: Did you complete all 3 requirements?
2. **AI Usage Quality (1-5)**: Based on your shared conversations/screenshots, did you use AI effectively?
3. **Code Integration (1-5)**: Does your code follow existing patterns and integrate well?

## Tips for Success

- **Read the existing code carefully** to understand the patterns
- **Use AI to help understand unfamiliar Go concepts** or syntax
- **Test your changes frequently** - run the server and test the endpoints
- **Follow the existing code style** and error handling patterns
- **Ask AI for help with Go-specific testing practices**
- **Don't hesitate to iterate** - refine your prompts if AI doesn't give you what you need

## Troubleshooting

- If you get import errors, run: `go mod tidy`
- If the server won't start, check that port 8080 is available
- If tests fail, run with verbose output: `go test -v`
- The application creates JSON files for data storage - this is normal

Good luck! Remember to document your AI usage throughout the process.
