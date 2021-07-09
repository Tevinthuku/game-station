package members

import (
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
)

type OnlineID string

type Member struct {
	OnlineID OnlineID
	SignInID entities.SignInID
}
