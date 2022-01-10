package repository

import (
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/util/logger"
)

type MemberInfoRepo interface {
	FindById(memberId string) (*entity.Member, error)
	Create(newMember *entity.Member) (*entity.Member, error)
}

type memberInfoRepo struct {
}

func (m *memberInfoRepo) FindById(memberId string) (*entity.Member, error) {
	logger.Log.Debug().Msgf("Find by %s", memberId)
	return nil, nil
}

func (m *memberInfoRepo) Create(newMember *entity.Member) (*entity.Member, error) {
	logger.Log.Debug().Msgf("Create %s", newMember)
	return nil, nil
}

func NewMemberPersonalInfoRepo() MemberInfoRepo {
	memberPersonalInfoRepo := &memberInfoRepo{}
	return memberPersonalInfoRepo
}
