package domain

import (
	"github.com/pkg/errors"

	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
)

type OnlineID string

type Member struct {
	OnlineID OnlineID
	SignInID entities.SignInID
}

var (
	ErrOnlineIDIsTaken            = errors.New("the onlineID is already in use")
	ErrSignInIDIsTaken            = errors.New("the signinID is already in use")
	ErrMemberWithOnlineIDNotFound = errors.New("the member with the onlineID isnt found")
	ErrMemberWithSignInIDNotFound = errors.New("the member with the signinID isnt found")

)
