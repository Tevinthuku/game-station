package members

import (
	networkentities "github.com/Tevinthuku/game-station/pkg/gamestationnetwork/accounts/entities"
	"github.com/Tevinthuku/game-station/pkg/gamestationplus/members/entities"
	"github.com/pkg/errors"
)

var (
	ErrOnlineIDIsTaken = errors.New("the onlineID is already in use")
	ErrSignInIDIsTaken = errors.New("the signinID is already in use")

	ErrMemberWithOnlineIDNotFound = errors.New("the member with the onlineID isnt found")
)

type (
	Repository interface {
		AddNewMember(newOnLineID entities.OnlineID, networkSignInID networkentities.SignInID) *entities.Member
		GetMemberByOnlineID(onlineID entities.OnlineID) (*entities.Member, error)
		GetMemberByNetworkSignInID(signinID networkentities.SignInID) (*entities.Member, error)
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

func (s *Service) JoinToPlayStationPlus(newOnLineID entities.OnlineID, networkSignInID networkentities.SignInID) (*entities.Member, error) {
	_, err := s.accountService.VerifyUserWithSignInIDExists(networkSignInID)
	if err != nil {
		return &entities.Member{}, err
	}
	err = s.verifyOnlineIDIsAvailable(newOnLineID)
	if err != nil {
		return &entities.Member{}, err
	}
	err = s.verifyNetworkSignInIDIsAvailable(networkSignInID)
	if err != nil {
		return &entities.Member{}, err
	}
	return s.memberRepo.AddNewMember(newOnLineID, networkSignInID), nil
}

func (s *Service) verifyOnlineIDIsAvailable(onlineID entities.OnlineID) error {
	_, err := s.memberRepo.GetMemberByOnlineID(onlineID)
	if err != nil {
		return nil
	}
	return ErrOnlineIDIsTaken
}

func (s *Service) verifyNetworkSignInIDIsAvailable(networkSignInID networkentities.SignInID) error {
	_, err := s.memberRepo.GetMemberByNetworkSignInID(networkSignInID)
	if err != nil {
		return nil
	}
	return ErrSignInIDIsTaken
}
