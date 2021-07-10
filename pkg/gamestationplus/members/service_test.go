package members

import (
	"testing"
	"time"

	"github.com/Tevinthuku/game-station/internal/storage/inmem"
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"
	networkDomain "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions"
	subscriptionDomain "github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMemberService(t *testing.T) {

	memberRepo := inmem.NewMembersStore()
	accountsRepo := inmem.NewAccountsStore()
	accountsService := accounts.NewService(accountsRepo)
	plusSubscriptionRepo := inmem.NewSubscriptionsStore()
	subscriptionService := subscriptions.NewService(plusSubscriptionRepo)
	service := NewService(memberRepo, accountsService, subscriptionService)
	subscriptionCode1 := subscriptionDomain.SubscriptionCode("FR2756290519")
	subscriptionCode2 := subscriptionDomain.SubscriptionCode("FR7971713142")
	subscriptionCode3 := subscriptionDomain.SubscriptionCode("CH6944356JHF")
	subscriptionCode4 := subscriptionDomain.SubscriptionCode("FR6183522024")
	acc1 := networkDomain.Account{
		SignInID:    networkDomain.SignInID("test@gamer.com"),
		UserName:    "testgamer",
		DateOfBirth: time.Date(1995, 12, 9, 0, 0, 0, 0, time.UTC), // TODO: Create a better abstraction for this
	}
	onlineID := domain.OnlineID("gamer")
	_, _ = accountsService.CreateAccount(acc1)
	Convey("Given an OnlineID", t, func() {
		Convey("it should let the account owner join gamestation plus if the SignInID is registered on gamestation network", func() {
			member, err := service.JoinToPlayStationPlus(onlineID, acc1.SignInID, subscriptionCode1)
			So(err, ShouldBeNil)
			So(member.OnlineID, ShouldEqual, onlineID)
		})
		Convey("it should not let the user join game station plus if the member's SignInID isnt registered on gamestation network", func() {
			onlineID := domain.OnlineID("gamer2")
			signInID := networkDomain.SignInID("test2@gmail.com")
			_, err := service.JoinToPlayStationPlus(onlineID, signInID, subscriptionCode2)
			So(err, ShouldNotBeNil)
		})
		Convey("it should not let an onlineID be used to join twice", func() {
			_, err := service.JoinToPlayStationPlus(onlineID, acc1.SignInID, subscriptionCode3)
			So(err, ShouldNotBeNil)
		})
		Convey("it should not allow a member to join if SignInID is already registered to a member", func() {
			onlineID := domain.OnlineID("gamer4")
			_, err := service.JoinToPlayStationPlus(onlineID, acc1.SignInID, subscriptionCode4)
			So(err, ShouldNotBeNil)
		})
	})
}
