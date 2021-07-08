package subscriptions

import (
	"time"

	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members"
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
		MemberID   members.OnlineID
	}
	CurrentMemberSubscription struct {
		ValidUntil time.Time
	}
)
