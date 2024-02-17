package models

import (
	"fmt"
	"regexp"
	"time"
)

// Member represents a member with their information
type Member struct {
	FullName    string            `json:"full_name"`
	Age         int               `json:"age"`
	Bio         string            `json:"bio"`
	Title       string            `json:"title"`
	SocialMedia map[string]string `json:"social_media"`
	DateJoined  time.Time         `json:"date_joined"`
}

// Validate validates the member information
func (m *Member) Validate() error {
	if m.FullName == "" {
		return fmt.Errorf("full name cannot be empty")
	}
	if m.Age <= 0 || m.Age < 18 {
		return fmt.Errorf("age must not be less than 18")
	}
	if m.Bio == "" {
		return fmt.Errorf("bio cannot be empty")
	}
	if m.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	for platform, url := range m.SocialMedia {
		// Basic URL validation example
		if !isValidURL(url) {
			return fmt.Errorf("invalid URL for platform %s: %s", platform, url)
		}
	}
	return nil
}

// isValidURL checks if the given string is a valid URL
func isValidURL(url string) bool {
	regex := regexp.MustCompile(`^(http|https)://[\w\-.]+\.[a-zA-Z]{2,6}/[^\s]+$`)
	return regex.MatchString(url)
}
