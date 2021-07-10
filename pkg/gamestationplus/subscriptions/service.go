package subscriptions

import (
	memberdomain "github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/domain"
)

type (
	Repository interface {
		AddSubscriptionToMember(subscription domain.Subscription, member memberdomain.Member) (*domain.Subscription, error)
		GetAllMemberSubscriptions(member *memberdomain.Member) []*domain.MemberSubscription
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

func (ss *Service) GetUnUsedSubscriptionFromCode(code domain.SubscriptionCode) (*domain.Subscription, error) {
	return ss.subscriptionRepo.GetUnUsedSubscriptionFromCode(code)
}

func (ss *Service) GetCurrentMemberSubscription(member *memberdomain.Member) (*domain.CurrentMemberSubscription, error) {
	memberSubscriptions := ss.subscriptionRepo.GetAllMemberSubscriptions(member)
	if len(memberSubscriptions) == 0 {
		return &domain.CurrentMemberSubscription{}, domain.ErrMemberHasNoSubscriptions
	}
	firstSubscription := memberSubscriptions[0]
	currentMemberSubscription := domain.CurrentMemberSubscription{
		MemberID:   member.OnlineID,
		ValidUntil: firstSubscription.ValidUntil(),
	}
	for i := range memberSubscriptions[1:] {
		if memberSubscriptions[i].IsBoughtBeforeExpiryOfCurrentSubscription(currentMemberSubscription) {
			currentMemberSubscription.ExtendWithDuration(memberSubscriptions[i].Duration)
		} else {
			currentMemberSubscription.ValidUntil = memberSubscriptions[i].ValidUntil()
		}
	}

	if currentMemberSubscription.IsExpired() {
		return &domain.CurrentMemberSubscription{}, domain.ErrCurrentSubscriptionIsExpired
	}
	return &currentMemberSubscription, nil
}
