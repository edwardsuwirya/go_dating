package usecase

import (
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/repository"
)

type MemberPreferenceUseCase interface {
	UpdatePreference(memberPref *entity.MemberPreferences) (*entity.MemberPreferences, error)
}
type memberPreferenceUseCase struct {
	prefRepo repository.MemberPreferenceRepo
}

func (m *memberPreferenceUseCase) UpdatePreference(memberPref *entity.MemberPreferences) (*entity.MemberPreferences, error) {
	return m.prefRepo.Create(memberPref)
}

func NewMemberPreferenceUseCase(repo repository.MemberPreferenceRepo) MemberPreferenceUseCase {
	return &memberPreferenceUseCase{prefRepo: repo}
}
