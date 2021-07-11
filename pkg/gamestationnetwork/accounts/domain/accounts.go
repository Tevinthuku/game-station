package domain

import (
	"time"

	"github.com/pkg/errors"
)

var (
	ErrSignInIdIsTaken            = errors.New("the SignInID is already in use")
	ErrUserNameIsTaken            = errors.New("the username is already in use")
	ErrNoAccountWithSignInIDFound = errors.New("no account found with specified signInID")
)

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
