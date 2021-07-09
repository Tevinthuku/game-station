package members

import (
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
)

type (
	Repository interface {
		AddNewMember(newOnLineID OnlineID, networkSignInID entities.SignInID) *Member
	}

	AccountService interface {
		VerifyUserWithSignInIDExists(signInID entities.SignInID) (*entities.Account, error)
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

func (s *Service) JoinToPlayStationPlus(newOnLineID OnlineID, networkSignInID entities.SignInID) (*Member, error) {
	_, err := s.accountService.VerifyUserWithSignInIDExists(networkSignInID)
	if err != nil {
		return &Member{}, err
	}
	return s.memberRepo.AddNewMember(newOnLineID, networkSignInID), nil
}
