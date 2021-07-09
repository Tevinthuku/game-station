package members

import (
	"testing"
	"time"

	"github.com/Tevinthuku/game-station/internal/storage/inmem"
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"
	networkentities "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/entities"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMemberService(t *testing.T) {

	memberRepo := inmem.NewMembersStore()
	accountsRepo := inmem.NewAccountsStore()
	accountsService := accounts.NewService(accountsRepo)
	service := NewService(memberRepo, accountsService)
	acc1 := networkentities.Account{
		SignInID:    networkentities.SignInID("test@gamer.com"),
		UserName:    "testgamer",
		DateOfBirth: time.Date(1995, 12, 9, 0, 0, 0, 0, time.UTC), // TODO: Create a better abstraction for this
	}
	onlineID := entities.OnlineID("gamer")
	_, _ = accountsService.CreateAccount(acc1)
	Convey("Given an OnlineID", t, func() {
		Convey("it should let the account owner join gamestation plus if the SignInID is registered on gamestation network", func() {
			member, err := service.JoinToPlayStationPlus(onlineID, acc1.SignInID)
			So(err, ShouldBeNil)
			So(member.OnlineID, ShouldEqual, onlineID)
		})
		Convey("it should not let the user join game station plus if the member's SignInID isnt registered on gamestation network", func() {
			onlineID := entities.OnlineID("gamer2")
			signInID := networkentities.SignInID("test2@gmail.com")
			_, err := service.JoinToPlayStationPlus(onlineID, signInID)
			So(err, ShouldNotBeNil)
		})
		Convey("it should not let an onlineID be used to join twice", func() {
			_, err := service.JoinToPlayStationPlus(onlineID, acc1.SignInID)
			So(err, ShouldNotBeNil)
		})
		Convey("it should not allow a member to join if SignInID is already registered to a member", func() {
			onlineID := entities.OnlineID("gamer4")
			_, err := service.JoinToPlayStationPlus(onlineID, acc1.SignInID)
			So(err, ShouldNotBeNil)
		})
	})
}
