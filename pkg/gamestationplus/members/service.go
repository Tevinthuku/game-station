package members

import (
	networkentities "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/entities"
)

type (
	Repository interface {
		AddNewMember(newOnLineID entities.OnlineID, networkSignInID networkentities.SignInID) *entities.Member
	}

	AccountService interface {
		VerifyUserWithSignInIDExists(signInID networkentities.SignInID) (*networkentities.Account, error)
	}

	Service struct {
		memberRepo     Repository
		accountService AccountService
	}
)

func NewService(memberRepo Repository, accountsService AccountService) *Service {
	return &Service{
		memberRepo:     memberRepo,
		accountService: accountsService,
	}
}

func (s *Service) JoinToPlayStationPlus(newOnLineID entities.OnlineID, networkSignInID networkentities.SignInID) (*entities.Member, error) {
	_, err := s.accountService.VerifyUserWithSignInIDExists(networkSignInID)
	if err != nil {
		return &entities.Member{}, err
	}
	return s.memberRepo.AddNewMember(newOnLineID, networkSignInID), nil
}
