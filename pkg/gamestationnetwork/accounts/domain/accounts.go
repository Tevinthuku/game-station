package domain

import "time"

type (
	DateOfBirth = time.Time
	Email       = string
	SignInID    (Email)
	Account     struct {
		SignInID    SignInID
		UserName    string
		DateOfBirth DateOfBirth
	}
)
