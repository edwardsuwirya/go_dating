package repository

import (
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/util/logger"
)

type MemberAccessRepo interface {
	FindByUserNameAndPassword(userName string, password string) (*entity.MemberUserAccess, error)
	UpdateVerification(id string) error
	Create(newAccess *entity.MemberUserAccess) error
}

type memberAccessRepo struct {
}

func (m *memberAccessRepo) FindByUserNameAndPassword(userName string, password string) (*entity.MemberUserAccess, error) {
	logger.Log.Debug().Msgf("Authenticate %s", userName)
	return nil, nil
}

func (m *memberAccessRepo) UpdateVerification(id string) error {
	logger.Log.Debug().Msgf("Update Verification %s", id)
	return nil
}

func (m *memberAccessRepo) Create(newAccess *entity.MemberUserAccess) error {
	logger.Log.Debug().Msgf("Create %s", newAccess.UserName)
	return nil
}

func NewMemberAccessRepo() MemberAccessRepo {
	memberAccessRepo := &memberAccessRepo{}
	return memberAccessRepo
}
