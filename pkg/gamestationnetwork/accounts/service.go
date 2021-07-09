package accounts

import (
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
	"github.com/pkg/errors"
)

var (
	ErrSignInIdIsTaken = errors.New("the SignInID is already in use")
	ErrUserNameIsTaken = errors.New("the username is already in use")

	ErrNoAccountWithSignInIDFound = errors.New("no account found with specified signInID")
)

type (
	Repository interface {
		AddNewAccount(accoount entities.Account) (*entities.Account, error)
		GetAccountBySignInId(signInID entities.SignInID) (*entities.Account, error)
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

func (s *Service) CreateAccount(account entities.Account) (*entities.Account, error) {
	return s.accountRepo.AddNewAccount(account)
}

func (s *Service) VerifyUserWithSignInIDExists(signInID entities.SignInID) (*entities.Account, error) {
	return s.accountRepo.GetAccountBySignInId(signInID)
}
