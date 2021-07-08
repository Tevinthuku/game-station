package inmem

import (
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members"
)

type MembersStorage struct {
	members []members.Member
}

func NewMembersStore() *MembersStorage {
	return &MembersStorage{}
}

func (ms *MembersStorage) AddNewMember(newOnLineID members.OnlineID, networkSignInID accounts.SignInID) *members.Member {
	member := members.Member{
		OnlineID: newOnLineID,
		SignInID: networkSignInID,
	}

	ms.members = append(ms.members, member)

	return &member
}
