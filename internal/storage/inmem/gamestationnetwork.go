package inmem

import (
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/domain"
)

type AccountsStorage struct {
	accounts []domain.Account
}

func NewAccountsStore() *AccountsStorage {
	return &AccountsStorage{}
}

func (as *AccountsStorage) AddNewAccount(account domain.Account) (*domain.Account, error) {
	for i := range as.accounts {
		if as.accounts[i].SignInID == account.SignInID {
			return &domain.Account{}, domain.ErrSignInIdIsTaken
		}
		if as.accounts[i].UserName == account.UserName {
			return &domain.Account{}, domain.ErrUserNameIsTaken
		}
	}
	as.accounts = append(as.accounts, account)
	return &account, nil
}

func (as *AccountsStorage) GetAccountBySignInId(signInID domain.SignInID) (*domain.Account, error) {
	for i := range as.accounts {
		if as.accounts[i].SignInID == signInID {
			return &as.accounts[i], nil
		}
	}
	return &domain.Account{}, domain.ErrNoAccountWithSignInIDFound
}
