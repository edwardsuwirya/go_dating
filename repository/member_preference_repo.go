package repository

import (
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/util/logger"
)

type MemberPreferenceRepo interface {
	Create(newMemberPref *entity.MemberPreferences) (*entity.MemberPreferences, error)
}

type memberPreferenceRepo struct {
}

func (m *memberPreferenceRepo) Create(newMemberPref *entity.MemberPreferences) (*entity.MemberPreferences, error) {
	logger.Log.Debug().Msgf("Create %s", newMemberPref)
	return nil, nil
}

func NewMemberPreferenceRepo() MemberPreferenceRepo {
	return &memberPreferenceRepo{}
}
