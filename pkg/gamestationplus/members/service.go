package members

import (
	networkentities "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/domain"
)

type (
	Repository interface {
		AddNewMember(newOnLineID domain.OnlineID, networkSignInID networkentities.SignInID) *domain.Member
		GetMemberByOnlineID(onlineID domain.OnlineID) (*domain.Member, error)
		GetMemberByNetworkSignInID(signinID networkentities.SignInID) (*domain.Member, error)
	}

	AccountService interface {
		VerifyUserWithSignInIDExists(signInID networkentities.SignInID) (*networkentities.Account, error)
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

func (s *Service) JoinToPlayStationPlus(newOnLineID domain.OnlineID, networkSignInID networkentities.SignInID) (*domain.Member, error) {
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

func (s *Service) verifyNetworkSignInIDIsAvailable(networkSignInID networkentities.SignInID) error {
	_, err := s.memberRepo.GetMemberByNetworkSignInID(networkSignInID)
	if err != nil {
		return nil
	}
	return domain.ErrSignInIDIsTaken
}
