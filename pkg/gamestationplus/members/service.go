package members

import (
	networkDomain "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/domain"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
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

	Service struct {
		memberRepo     Repository
		accountService AccountService
	}
)

func NewService(memberRepo Repository, accountsService AccountService) *Service {
	return &Service{
		memberRepo:     memberRepo,
		accountService: accountsService,
	}
}

func (s *Service) JoinToPlayStationPlus(newOnLineID domain.OnlineID, networkSignInID networkDomain.SignInID) (*domain.Member, error) {
	_, err := s.accountService.VerifyUserWithSignInIDExists(networkSignInID)
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
	return s.memberRepo.AddNewMember(newOnLineID, networkSignInID), nil
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
