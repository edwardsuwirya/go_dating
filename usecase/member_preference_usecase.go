package usecase

import (
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/repository"
)

type MemberPreferenceUseCase interface {
	CreatePreference(memberPref *entity.MemberPreferences, memberInterest []entity.MemberInterest) error
}
type memberPreferenceUseCase struct {
	prefRepo repository.MemberPreferenceRepo
}

func (m *memberPreferenceUseCase) CreatePreference(memberPref *entity.MemberPreferences, memberInterest []entity.MemberInterest) error {
	return m.prefRepo.Create(memberPref, memberInterest)
}

func NewMemberPreferenceUseCase(repo repository.MemberPreferenceRepo) MemberPreferenceUseCase {
	return &memberPreferenceUseCase{prefRepo: repo}
}
