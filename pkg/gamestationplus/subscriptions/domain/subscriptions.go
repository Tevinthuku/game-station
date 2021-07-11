package domain

import (
	"time"

	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	"github.com/pkg/errors"
)

var (
	ErrNoSubscriptionWithCodeFound  = errors.New("no subscription found with specified subscription code")
	ErrMemberHasNoSubscriptions     = errors.New("member does not have any subscription")
	ErrCurrentSubscriptionIsExpired = errors.New("the member has no active subscriptions")
)

type Duration (string)

const (
	ONEMONTH    Duration = "ONEMONTH"
	THREEMONTHS Duration = "THREEMONTHS"
	SIXMONTHS   Duration = "SIXMONTHS"
	ONEYEAR     Duration = "ONEYEAR"
)

type (
	SubscriptionCode (string)
	Subscription     struct {
		Duration Duration
		Code     SubscriptionCode
	}
	MemberSubscription struct {
		Code       SubscriptionCode
		DateBought time.Time
		Duration   Duration
		MemberID   domain.OnlineID
	}
	CurrentMemberSubscription struct {
		MemberID   domain.OnlineID
		ValidUntil time.Time
	}
)

func (cs *CurrentMemberSubscription) IsExpired() bool {
	return time.Now().After(cs.ValidUntil)
}

func (cs *CurrentMemberSubscription) ExtendWithDuration(duration Duration) {
	cs.ValidUntil = extendDateWithDuration(cs.ValidUntil, duration)
}

func (ms *MemberSubscription) ValidUntil() time.Time {
	return extendDateWithDuration(ms.DateBought, ms.Duration)
}

func (ms *MemberSubscription) IsBoughtBeforeExpiryOfCurrentSubscription(currentSubscription CurrentMemberSubscription) bool {
	return ms.DateBought.Before(currentSubscription.ValidUntil)
}

func extendDateWithDuration(date time.Time, duration Duration) time.Time {
	switch duration {
	case ONEMONTH:
		return date.AddDate(0, 1, 0)
	case THREEMONTHS:
		return date.AddDate(0, 3, 0)
	case SIXMONTHS:
		return date.AddDate(0, 6, 0)
	case ONEYEAR:
		return date.AddDate(1, 0, 0)
	default:
		return date
	}
}
