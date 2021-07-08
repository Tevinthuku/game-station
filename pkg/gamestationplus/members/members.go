package members

import "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"

type OnlineID string

type Member struct {
	OnlineID OnlineID
	SignInID accounts.SignInID
}
