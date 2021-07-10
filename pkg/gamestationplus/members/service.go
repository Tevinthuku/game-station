package members

import (
	networkDomain "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
	subscriptionDomain "github.com/Tevinthuku/game-station/pkg/gamestationplus/subscriptions/domain"
)

type (
	Repository interface {
		AddNewMember(newOnLineID domain.OnlineID, networkSignInID networkDomain.SignInID) *domain.Member
		GetMemberByOnlineID(onlineID domain.OnlineID) (*domain.Member, error)
		GetMemberByNetworkSignInID(signinID networkDomain.SignInID) (*domain.Member, error)
	}

	AccountService interface {
		VerifyUserWithSignInIDExists(signInID networkDomain.SignInID) (*networkDomain.Account, error)
	}

	SubscriptionService interface {
		GetCurrentMemberSubscription(member *domain.Member) (*subscriptionDomain.CurrentMemberSubscription, error)
		VerifySubscriptionCodeIsValid(code subscriptionDomain.SubscriptionCode) (*subscriptionDomain.Subscription, error)
		AddSubscriptionToMember(subscription *subscriptionDomain.Subscription, member *domain.Member)
	}

	Service struct {
		memberRepo          Repository
		accountService      AccountService
		subscriptionService SubscriptionService
	}
)

func NewService(memberRepo Repository, accountsService AccountService, subscriptionService SubscriptionService) *Service {
	return &Service{
		memberRepo:          memberRepo,
		accountService:      accountsService,
		subscriptionService: subscriptionService,
	}
}

func (s *Service) JoinToPlayStationPlus(newOnLineID domain.OnlineID, networkSignInID networkDomain.SignInID, subscriptionCode subscriptionDomain.SubscriptionCode) (*domain.Member, error) {
	subscription, err := s.subscriptionService.VerifySubscriptionCodeIsValid(subscriptionCode)
	if err != nil {
		return &domain.Member{}, err
	}
	_, err = s.accountService.VerifyUserWithSignInIDExists(networkSignInID)
	if err != nil {
		return &domain.Member{}, err
	}
	err = s.verifyOnlineIDIsAvailable(newOnLineID)
	if err != nil {
		return &domain.Member{}, err
	}
	err = s.verifyNetworkSignInIDIsAvailable(networkSignInID)
	if err != nil {
		return &domain.Member{}, err
	}
	member := s.memberRepo.AddNewMember(newOnLineID, networkSignInID)

	s.subscriptionService.AddSubscriptionToMember(subscription, member)

	return member, nil

}

func (s *Service) verifyOnlineIDIsAvailable(onlineID domain.OnlineID) error {
	_, err := s.memberRepo.GetMemberByOnlineID(onlineID)
	if err != nil {
		return nil
	}
	return domain.ErrOnlineIDIsTaken
}

func (s *Service) verifyNetworkSignInIDIsAvailable(networkSignInID networkDomain.SignInID) error {
	_, err := s.memberRepo.GetMemberByNetworkSignInID(networkSignInID)
	if err != nil {
		return nil
	}
	return domain.ErrSignInIDIsTaken
}

func (s *Service) GetCurrentMemberSubscription(onlineID domain.OnlineID) (*subscriptionDomain.CurrentMemberSubscription, error) {
	member, err := s.memberRepo.GetMemberByOnlineID(onlineID)
	if err != nil {
		return &subscriptionDomain.CurrentMemberSubscription{}, err
	}
	return s.subscriptionService.GetCurrentMemberSubscription(member)
}
