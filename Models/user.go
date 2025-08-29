// File: models/user.go
//This file is an example of what to achieve using standard Go practices
// Package models contains the data structures for our application.
package models

import (
	"errors"
	"strings"
	"time"
)

// User represents a user in the system.
// The struct tags define how this struct will be encoded to and from JSON.
type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // The '-' tag prevents this field from being sent in JSON responses.
	FirstName    string    `json:"firstName,omitempty"`
	LastName     string    `json:"lastName,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
}

// NewUser is a constructor function to create a new User instance.
// It validates the input and returns an error if the data is invalid.
func NewUser(email, password, firstName, lastName string) (*User, error) {
	// Basic validation
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	// In a real application, you would hash the password here.
	// For this example, we'll store it directly, but this is NOT secure.
	passwordHash := "hashed_" + password // Placeholder for a real hash

	return &User{
		// ID would typically be set by the database, so we leave it as 0.
		Email:        strings.ToLower(email), // Store email in lowercase for consistency.
		PasswordHash: passwordHash,
		FirstName:    firstName,
		LastName:     lastName,
		CreatedAt:    time.Now().UTC(), // Use UTC for consistency across timezones.
	}, nil
}

// FullName is a method on the User struct.
// It returns the user's full name by combining FirstName and LastName.
func (u *User) FullName() string {
	// Trim space to handle cases where one name might be missing.
	return strings.TrimSpace(u.FirstName + " " + u.LastName)
}
