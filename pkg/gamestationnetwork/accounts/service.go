package accounts

import (
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/domain"
	"github.com/pkg/errors"
)

var (
	ErrSignInIdIsTaken = errors.New("the SignInID is already in use")
	ErrUserNameIsTaken = errors.New("the username is already in use")

	ErrNoAccountWithSignInIDFound = errors.New("no account found with specified signInID")
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
