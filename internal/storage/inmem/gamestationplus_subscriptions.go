package inmem

import (
	"time"

	memberdomain "github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/domain"
)

type storedAvailableSubscription struct {
	domain.Subscription
	isAvailable bool
}
type SubscriptionStore struct {
	membersubscriptions []domain.MemberSubscription
	allSubscriptions    []storedAvailableSubscription
}

func NewSubscriptionsStore() *SubscriptionStore {
	return &SubscriptionStore{}
}

func (ss *SubscriptionStore) AddSubscriptionToMember(subscription domain.Subscription, member memberdomain.Member) (*domain.Subscription, error) {
	memberSubscription := domain.MemberSubscription{
		Code:       subscription.Code,
		DateBought: time.Now(),
		Duration:   subscription.Duration,
	}

	for i := range ss.allSubscriptions {
		if ss.allSubscriptions[i].isAvailable && ss.allSubscriptions[i].Code == subscription.Code {
			ss.membersubscriptions = append(ss.membersubscriptions, memberSubscription)
			ss.allSubscriptions[i].isAvailable = false
			return &subscription, nil
		}
	}

	return &domain.Subscription{}, domain.ErrNoSubscriptionWithCodeFound
}

func (ss *SubscriptionStore) GetAllMemberSubscriptions(member memberdomain.Member) []*domain.MemberSubscription {
	memberSubscriptions := []*domain.MemberSubscription{}
	for i := range ss.membersubscriptions {
		if ss.membersubscriptions[i].MemberID == member.OnlineID {
			memberSubscriptions = append(memberSubscriptions, &ss.membersubscriptions[i])
		}
	}
	return memberSubscriptions
}

func (ss *SubscriptionStore) GetUnUsedSubscriptionFromCode(code domain.SubscriptionCode) (*domain.Subscription, error) {
	for i := range ss.allSubscriptions {
		if ss.allSubscriptions[i].isAvailable && ss.allSubscriptions[i].Code == code {
			return &ss.allSubscriptions[i].Subscription, nil
		}
	}
	return &domain.Subscription{}, domain.ErrNoSubscriptionWithCodeFound
}
