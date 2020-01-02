// Package users provides support for user management.
package users

// user represents information about a user.
type user struct {
	Name string
	ID   int
}

// Manager represents information about a manager.
type Manager struct {
	Title string

	user
}
