package inmem

import (
	networkDomain "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
)

type MembersStorage struct {
	members []domain.Member
}

func NewMembersStore() *MembersStorage {
	return &MembersStorage{}
}

func (ms *MembersStorage) AddNewMember(newOnLineID domain.OnlineID, networkSignInID networkDomain.SignInID) *domain.Member {
	member := domain.Member{
		OnlineID: newOnLineID,
		SignInID: networkSignInID,
	}

	ms.members = append(ms.members, member)

	return &member
}

func (ms *MembersStorage) GetMemberByOnlineID(onlineID domain.OnlineID) (*domain.Member, error) {
	for i := range ms.members {
		if ms.members[i].OnlineID == onlineID {
			return &ms.members[i], nil
		}
	}
	return &domain.Member{}, domain.ErrMemberWithOnlineIDNotFound
}

func (ms *MembersStorage) GetMemberByNetworkSignInID(signInID networkDomain.SignInID) (*domain.Member, error) {
	for i := range ms.members {
		if ms.members[i].SignInID == signInID {
			return &ms.members[i], nil
		}
	}
	return &domain.Member{}, domain.ErrMemberWithSignInIDNotFound
}
