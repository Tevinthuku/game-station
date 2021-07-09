package domain

import (
	"time"

	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	"github.com/pkg/errors"
)

var (
	ErrNoSubscriptionWithCodeFound = errors.New("no subscription found with specified subscription code")
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
		ValidUntil time.Time
	}
)
