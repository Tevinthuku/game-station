package accounts

import (
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/domain"
)


type (
	Repository interface {
		AddNewAccount(accoount domain.Account) (*domain.Account, error)
		GetAccountBySignInId(signInID domain.SignInID) (*domain.Account, error)
	}

	Service struct {
		accountRepo Repository
	}
)

func NewService(accountRepo Repository) *Service {
	return &Service{
		accountRepo,
	}
}

func (s *Service) CreateAccount(account domain.Account) (*domain.Account, error) {
	return s.accountRepo.AddNewAccount(account)
}

func (s *Service) VerifyUserWithSignInIDExists(signInID domain.SignInID) (*domain.Account, error) {
	return s.accountRepo.GetAccountBySignInId(signInID)
}
