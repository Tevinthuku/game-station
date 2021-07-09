package main

import (
	"github.com/Tevinthuku/game-station/internal/storage/inmem"
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions"
)

func main() {
	accountsRepo := inmem.NewAccountsStore()
	plusMemberRepo := inmem.NewMembersStore()
	plusSubscriptionRepo := inmem.NewSubscriptionsStore()

	accountsService := accounts.NewService(accountsRepo)
	_ = members.NewService(plusMemberRepo, accountsService)

	_ = subscriptions.NewService(plusSubscriptionRepo)

}
