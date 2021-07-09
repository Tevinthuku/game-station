package members

import "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"

type (
	Repository interface {
		AddNewMember(newOnLineID OnlineID, networkSignInID accounts.SignInID) *Member
	}

	AccountService interface {
		VerifyUserWithSignInIDExists(signInID accounts.SignInID) error
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

func (s *Service) JoinToPlayStationPlus(newOnLineID OnlineID, networkSignInID accounts.SignInID) (*Member, error) {
	err := s.accountService.VerifyUserWithSignInIDExists(networkSignInID)
	if err != nil {
		return &Member{}, err
	}
	return s.memberRepo.AddNewMember(newOnLineID, networkSignInID), nil
}
