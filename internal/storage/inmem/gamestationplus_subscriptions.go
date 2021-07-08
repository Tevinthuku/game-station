package inmem

import (
	"time"

	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions"
)

type storedAvailableSubscription struct {
	subscriptions.Subscription
	isAvailable bool
}
type SubscriptionStore struct {
	membersubscriptions []subscriptions.MemberSubscription
	allSubscriptions    []storedAvailableSubscription
}

func NewSubscriptionsStore() *SubscriptionStore {
	return &SubscriptionStore{}
}

func (ss *SubscriptionStore) AddSubscriptionToMember(subscription subscriptions.Subscription, member members.Member) (*subscriptions.Subscription, error) {
	memberSubscription := subscriptions.MemberSubscription{
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

	return &subscriptions.Subscription{}, subscriptions.ErrNoSubscriptionWithCodeFound
}

func (ss *SubscriptionStore) GetAllMemberSubscriptions(member members.Member) []*subscriptions.MemberSubscription {
	memberSubscriptions := []*subscriptions.MemberSubscription{}
	for i := range ss.membersubscriptions {
		if ss.membersubscriptions[i].MemberID == member.OnlineID {
			memberSubscriptions = append(memberSubscriptions, &ss.membersubscriptions[i])
		}
	}
	return memberSubscriptions
}

func (ss *SubscriptionStore) GetUnUsedSubscriptionFromCode(code subscriptions.SubscriptionCode) (*subscriptions.Subscription, error) {
	for i := range ss.allSubscriptions {
		if ss.allSubscriptions[i].isAvailable && ss.allSubscriptions[i].Code == code {
			return &ss.allSubscriptions[i].Subscription, nil
		}
	}
	return &subscriptions.Subscription{}, subscriptions.ErrNoSubscriptionWithCodeFound
}
