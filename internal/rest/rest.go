package rest

import (
	"github.com/Tevinthuku/game-station/internal/storage/inmem"
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions"
)

// TODO: Will use a http library to expose the required business rules via rest.
func StartServer() {
	accountsRepo := inmem.NewAccountsStore()
	plusMembersRepo := inmem.NewMembersStore()
	plusSubscriptionRepo := inmem.NewSubscriptionsStore()

	subscriptionService := subscriptions.NewService(plusSubscriptionRepo)
	accountsService := accounts.NewService(accountsRepo)
	_ = members.NewService(plusMembersRepo, accountsService, subscriptionService)

}
