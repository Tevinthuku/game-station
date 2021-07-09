package subscriptions

import (
	memberentities "github.com/Tevinthuku/game-station/pkg/gamestationplus/members/entities"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/entities"
	"github.com/pkg/errors"
)

var (
	ErrNoSubscriptionWithCodeFound = errors.New("no subscription found with specified subscription code")
)

type (
	Repository interface {
		AddSubscriptionToMember(subscription entities.Subscription, member memberentities.Member) (*entities.Subscription, error)
		GetAllMemberSubscriptions(member memberentities.Member) []*entities.MemberSubscription
		GetUnUsedSubscriptionFromCode(code entities.SubscriptionCode) (*entities.Subscription, error)
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
