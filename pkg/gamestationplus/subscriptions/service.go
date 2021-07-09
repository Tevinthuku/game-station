package subscriptions

import (
	memberdomain "github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/domain"
	"github.com/pkg/errors"
)

var (
	ErrNoSubscriptionWithCodeFound = errors.New("no subscription found with specified subscription code")
)

type (
	Repository interface {
		AddSubscriptionToMember(subscription domain.Subscription, member memberdomain.Member) (*domain.Subscription, error)
		GetAllMemberSubscriptions(member memberdomain.Member) []*domain.MemberSubscription
		GetUnUsedSubscriptionFromCode(code domain.SubscriptionCode) (*domain.Subscription, error)
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
