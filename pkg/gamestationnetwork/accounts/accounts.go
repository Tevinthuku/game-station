package accounts

import "time"

type (
	DateOfBirth = time.Duration
	Email       = string
	SignInID    (Email)
	Account     struct {
		SignInID    SignInID
		UserName    string
		DateOfBirth DateOfBirth
	}
)
