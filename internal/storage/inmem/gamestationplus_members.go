package inmem

import (
	networkentities "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/entities"
	"github.com/pkg/errors"
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

func (ms *MembersStorage) GetMemberByOnlineID(onlineID entities.OnlineID) (*entities.Member, error) {
	for i := range ms.members {
		if ms.members[i].OnlineID == onlineID {
			return &ms.members[i], nil
		}
	}
	return &entities.Member{}, errors.New("couldnt find user with onlineID")
}

func (ms *MembersStorage) GetMemberByNetworkSignInID(signInID networkentities.SignInID) (*entities.Member, error) {
	for i := range ms.members {
		if ms.members[i].SignInID == signInID {
			return &ms.members[i], nil
		}
	}
	return &entities.Member{}, errors.New("couldnt find user with signInID")
}
