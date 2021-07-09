package inmem

import (
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
)

type AccountsStorage struct {
	accounts []entities.Account
}

func NewAccountsStore() *AccountsStorage {
	return &AccountsStorage{}
}

func (as *AccountsStorage) AddNewAccount(account entities.Account) (*entities.Account, error) {
	for i := range as.accounts {
		if as.accounts[i].SignInID == account.SignInID {
			return &entities.Account{}, accounts.ErrSignInIdIsTaken
		}
		if as.accounts[i].UserName == account.UserName {
			return &entities.Account{}, accounts.ErrUserNameIsTaken
		}
	}
	as.accounts = append(as.accounts, account)
	return &account, nil
}

func (as *AccountsStorage) GetAccountBySignInId(signInID entities.SignInID) (*entities.Account, error) {
	for i := range as.accounts {
		if as.accounts[i].SignInID == signInID {
			return &as.accounts[i], nil
		}
	}
	return &entities.Account{}, accounts.ErrNoAccountWithSignInIDFound
}
