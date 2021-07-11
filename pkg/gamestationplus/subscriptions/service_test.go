package subscriptions

import (
	"testing"
	"time"

	networkDomain "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/domain"
	membersdomain "github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/domain"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/Tevinthuku/game-station/internal/storage/inmem"
	"github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members"
)

func getListOfValidSubscriptionCodes(count int) []domain.SubscriptionCode {
	subscriptions := inmem.PopulateSubscriptions()
	codes := make([]domain.SubscriptionCode, count)
	for i := 0; i < count; i++ {
		codes[i] = subscriptions[i].Code
	}
	return codes
}

type InactiveMemberMock struct{}

func NewInactiveMemberMockRepo() *InactiveMemberMock {
	return &InactiveMemberMock{}
}

func (ss *InactiveMemberMock) AddSubscriptionToMember(subscription *domain.Subscription, member *membersdomain.Member) {
}

func (ss *InactiveMemberMock) GetAllMemberSubscriptions(member *membersdomain.Member) []*domain.MemberSubscription {
	memberSubscriptions := []*domain.MemberSubscription{{
		Code:       domain.SubscriptionCode("EXPIRED"),
		Duration:   domain.ONEMONTH,
		DateBought: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
	}, {
		Code:       domain.SubscriptionCode("EXPIRED"),
		Duration:   domain.ONEMONTH,
		DateBought: time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC),
	}}
	return memberSubscriptions
}

func (ss *InactiveMemberMock) GetUnUsedSubscriptionFromCode(_code domain.SubscriptionCode) (*domain.Subscription, error) {

	return &domain.Subscription{}, nil
}

func TestSubscriptionService(t *testing.T) {
	memberRepo := inmem.NewMembersStore()
	accountsRepo := inmem.NewAccountsStore()

	accountsService := accounts.NewService(accountsRepo)
	plusSubscriptionRepo := inmem.NewSubscriptionsStore()
	subscriptionService := NewService(plusSubscriptionRepo)
	memberService := members.NewService(memberRepo, accountsService, subscriptionService)
	subscriptionCodes := getListOfValidSubscriptionCodes(5)
	acc1 := networkDomain.Account{
		SignInID:    networkDomain.SignInID("test@gamer.com"),
		UserName:    "testgamer",
		DateOfBirth: time.Date(1995, 12, 9, 0, 0, 0, 0, time.UTC), // TODO: Create a better abstraction for this
	}
	onlineID := membersdomain.OnlineID("gamer")
	_, _ = accountsService.CreateAccount(acc1)
	_, _ = memberService.JoinToPlayStationPlus(onlineID, acc1.SignInID, subscriptionCodes[0])
	Convey("Given a subscription and a member", t, func() {
		Convey("an already used code should not be available for use again", func() {
			_, err := subscriptionService.GetUnUsedSubscriptionFromCode(subscriptionCodes[0])
			So(err, ShouldNotBeNil)
		})
		Convey("a current member subscription should be valid if validity is past right now", func() {
			currentSubscription, err := memberService.GetCurrentMemberSubscription(onlineID)
			So(err, ShouldBeNil)
			So(currentSubscription.IsExpired(), ShouldBeFalse)
		})
		Convey("given an in-active member", func() {
			subscriptionRepo := NewInactiveMemberMockRepo()
			service := NewService(subscriptionRepo)
			Convey("their current subscription should be expired", func() {
				m := &membersdomain.Member{}
				currentSubscription, err := service.GetCurrentMemberSubscription(m)
				So(err, ShouldNotBeNil)
				So(currentSubscription.IsExpired(), ShouldBeTrue)
			})
		})
	})
}
