package models

// User represents a user in the application
type User struct {
	ID       string
	Email    string
	Name     string
	Picture  string
	Provider string
}

// NewUser creates a new User instance
func NewUser(id, email, name, picture, provider string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Name:     name,
		Picture:  picture,
		Provider: provider,
	}
}
