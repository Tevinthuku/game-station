package accounts

import "github.com/pkg/errors"

var (
	ErrSignInIdIsTaken = errors.New("the SignInID is already in use")
	ErrUserNameIsTaken = errors.New("the username is already in use")

	ErrNoAccountWithSignInIDFound = errors.New("no account found with specified signInID")
)

type (
	Repository interface {
		AddNewAccount(accoount Account) (*Account, error)
		GetAccountBySignInId(signInID SignInID) (*Account, error)
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

func (s *Service) CreateAccount(account Account) (*Account, error) {
	return s.accountRepo.AddNewAccount(account)
}

func (s *Service) GetAccountBySignInID(signinID SignInID) (*Account, error) {
	return s.accountRepo.GetAccountBySignInId(signinID)
}

func (s *Service) VerifyUserWithSignInIDExists(signInID SignInID) error {
	_, err := s.GetAccountBySignInID(signInID)
	return err
}
