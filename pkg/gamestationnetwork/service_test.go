package gamestationnetwork

import (
	"testing"
	"time"

	"github.com/Tevinthuku/game-station/internal/storage/inmem"
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAccountsService(t *testing.T) {
	accountsRepo := inmem.NewAccountsStore()
	service := accounts.NewService(accountsRepo)
	acc1 := entities.Account{
		SignInID:    entities.SignInID("test@gamer.com"),
		UserName:    "testgamer",
		DateOfBirth: time.Date(1995, 12, 9, 0, 0, 0, 0, time.UTC), // TODO: Create a better abstraction for this
	}
	Convey("Given an account", t, func() {
		Convey("it should be registered successfully", func() {
			acc, err := service.CreateAccount(acc1)
			So(err, ShouldBeNil)
			So(acc.SignInID, ShouldEqual, acc1.SignInID)
		})
		Convey("registration of same account again should fail", func() {
			_, err := service.CreateAccount(acc1)
			So(err, ShouldNotBeNil)
		})
	})
	Convey("Given a SignInID", t, func() {
		Convey("I should get an account that matches the SignInID if it exists", func() {
			acc, err := service.VerifyUserWithSignInIDExists(acc1.SignInID)
			So(err, ShouldBeNil)
			So(acc.SignInID, ShouldEqual, acc1.SignInID)
		})
		Convey("I should get an error if no account exists with the SignInID", func() {
			_, err := service.VerifyUserWithSignInIDExists(entities.SignInID("idont@exist.com"))
			So(err, ShouldNotBeNil)
		})
	})
}
