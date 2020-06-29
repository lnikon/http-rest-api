package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

// BeforeCreate ...
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
	}

	return nil
}

func encryptString(s string) (string, string) {
	b, err := bcrypt.GenerateFromPassword[]byte(s), bbcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}