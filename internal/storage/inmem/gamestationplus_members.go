package inmem

import (
	networkentities "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/entities"
)

type MembersStorage struct {
	members []entities.Member
}

func NewMembersStore() *MembersStorage {
	return &MembersStorage{}
}

func (ms *MembersStorage) AddNewMember(newOnLineID entities.OnlineID, networkSignInID networkentities.SignInID) *entities.Member {
	member := entities.Member{
		OnlineID: newOnLineID,
		SignInID: networkSignInID,
	}

	ms.members = append(ms.members, member)

	return &member
}
