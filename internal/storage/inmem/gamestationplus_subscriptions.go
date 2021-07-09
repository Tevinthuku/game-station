package inmem

import (
	"time"

	memberentities "github.com/Tevinthuku/game-station/pkg/gamestationplus/members/entities"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/entities"
)

type storedAvailableSubscription struct {
	entities.Subscription
	isAvailable bool
}
type SubscriptionStore struct {
	membersubscriptions []entities.MemberSubscription
	allSubscriptions    []storedAvailableSubscription
}

func NewSubscriptionsStore() *SubscriptionStore {
	return &SubscriptionStore{}
}

func (ss *SubscriptionStore) AddSubscriptionToMember(subscription entities.Subscription, member memberentities.Member) (*entities.Subscription, error) {
	memberSubscription := entities.MemberSubscription{
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

	return &entities.Subscription{}, subscriptions.ErrNoSubscriptionWithCodeFound
}

func (ss *SubscriptionStore) GetAllMemberSubscriptions(member memberentities.Member) []*entities.MemberSubscription {
	memberSubscriptions := []*entities.MemberSubscription{}
	for i := range ss.membersubscriptions {
		if ss.membersubscriptions[i].MemberID == member.OnlineID {
			memberSubscriptions = append(memberSubscriptions, &ss.membersubscriptions[i])
		}
	}
	return memberSubscriptions
}

func (ss *SubscriptionStore) GetUnUsedSubscriptionFromCode(code entities.SubscriptionCode) (*entities.Subscription, error) {
	for i := range ss.allSubscriptions {
		if ss.allSubscriptions[i].isAvailable && ss.allSubscriptions[i].Code == code {
			return &ss.allSubscriptions[i].Subscription, nil
		}
	}
	return &entities.Subscription{}, subscriptions.ErrNoSubscriptionWithCodeFound
}
