package subscriptions

import (
	memberdomain "github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/domain"
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

func (ss *Service) AddSubscriptionToMember(subscription domain.Subscription, member memberdomain.Member) (*domain.Subscription, error) {
	return ss.subscriptionRepo.AddSubscriptionToMember(subscription, member)
}

func (ss *Service) GetAllMemberSubscriptions(member memberdomain.Member) []*domain.MemberSubscription {
	return ss.subscriptionRepo.GetAllMemberSubscriptions(member)
}

func (ss *Service) GetUnUsedSubscriptionFromCode(code domain.SubscriptionCode) (*domain.Subscription, error) {
	return ss.subscriptionRepo.GetUnUsedSubscriptionFromCode(code)
}
