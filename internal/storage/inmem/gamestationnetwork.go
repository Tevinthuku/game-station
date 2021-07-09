package inmem

import "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"

type AccountsStorage struct {
	accounts []accounts.Account
}

func NewAccountsStore() *AccountsStorage {
	return &AccountsStorage{}
}

func (as *AccountsStorage) AddNewAccount(account accounts.Account) (*accounts.Account, error) {
	for i := range as.accounts {
		if as.accounts[i].SignInID == account.SignInID {
			return &accounts.Account{}, accounts.ErrSignInIdIsTaken
		}
		if as.accounts[i].UserName == account.UserName {
			return &accounts.Account{}, accounts.ErrUserNameIsTaken
		}
	}
	as.accounts = append(as.accounts, account)
	return &account, nil
}

func (as *AccountsStorage) GetAccountBySignInId(signInID accounts.SignInID) (*accounts.Account, error) {
	for i := range as.accounts {
		if as.accounts[i].SignInID == signInID {
			return &as.accounts[i], nil
		}
	}
	return &accounts.Account{}, accounts.ErrNoAccountWithSignInIDFound
}
