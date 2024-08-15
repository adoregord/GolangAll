package gosolid

import "github.com/gofrs/uuid"

// In this case, the User struct has two responsibilities: managing user data and saving it to the database.
// To adhere to the Single Responsibility Principle

type User struct {
	FirstName string
	LastName  string
	ID        uuid.UUID
	sources   string
}

//before srp impementation

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Save() error {
	// Save user to the database
	// ...
	return nil
}

//after srp impementation

// managing user data
func (u *User) GetFullName1() string {
	return u.FirstName + " " + u.LastName
}

type UserRepository struct {
	// Database connection or other storage-related fields
}

// saving to db
func (r *UserRepository) Save(u *User) error {
	// Save user to the database
	return nil
}

// Now, the User struct is only responsible for managing user data, while the UserRepository handles database operations
