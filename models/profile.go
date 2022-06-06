package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Profile struct {
	ID       int
	Login    string `json:"login" gorm:"unique"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      uint   `json:"age"`
}

// HashPassword hashes provided string
func (profile *Profile) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	profile.Password = string(bytes)
	return nil
}

// CheckPassword checks if provided string is the same as the one in the database
func (profile *Profile) CheckPassword(givenPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(givenPassword))
	if err != nil {
		return err
	}

	return nil
}
