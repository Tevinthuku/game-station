package subscriptions

import (
	memberdomain "github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/entities"
	"github.com/pkg/errors"
)

var (
	ErrNoSubscriptionWithCodeFound = errors.New("no subscription found with specified subscription code")
)

type (
	Repository interface {
		AddSubscriptionToMember(subscription entities.Subscription, member memberdomain.Member) (*entities.Subscription, error)
		GetAllMemberSubscriptions(member memberdomain.Member) []*entities.MemberSubscription
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
