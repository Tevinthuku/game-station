package inmem

import (
	"time"

	memberdomain "github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/domain"
)

type StoredAvailableSubscription struct {
	domain.Subscription
	isAvailable bool
}

type SubscriptionStore struct {
	membersubscriptions []domain.MemberSubscription
	allSubscriptions    []StoredAvailableSubscription
}

func NewSubscriptionsStore() *SubscriptionStore {
	return &SubscriptionStore{
		allSubscriptions: PopulateSubscriptions(),
	}
}

func (ss *SubscriptionStore) AddSubscriptionToMember(subscription *domain.Subscription, member *memberdomain.Member) {
	memberSubscription := domain.MemberSubscription{
		Code:       subscription.Code,
		DateBought: time.Now(),
		Duration:   subscription.Duration,
	}

	for i := range ss.allSubscriptions {
		if ss.allSubscriptions[i].Code == subscription.Code {
			ss.allSubscriptions[i].isAvailable = false
		}
	}
	ss.membersubscriptions = append(ss.membersubscriptions, memberSubscription)

}

func (ss *SubscriptionStore) GetAllMemberSubscriptions(member *memberdomain.Member) []*domain.MemberSubscription {
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
