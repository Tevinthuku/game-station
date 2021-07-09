package subscriptions

import (
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/entities"
	"github.com/pkg/errors"
)

var (
	ErrNoSubscriptionWithCodeFound = errors.New("no subscription found with specified subscription code")
)

type (
	Repository interface {
		AddSubscriptionToMember(subscription Subscription, member entities.Member) (*Subscription, error)
		GetAllMemberSubscriptions(member entities.Member) []*MemberSubscription
		GetUnUsedSubscriptionFromCode(code SubscriptionCode) (*Subscription, error)
	}

	Service struct {
		subscriptionRepo Repository
	}
)

func NewService(subscriptionRepo Repository) *Service {
	return &Service{
		subscriptionRepo: subscriptionRepo,
	}
}
