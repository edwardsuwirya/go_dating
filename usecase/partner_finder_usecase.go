package usecase

import (
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/repository"
)

type PartnerFinderUseCase interface {
	ViewPartner(id string, pageNo int, pageSize int) ([]entity.MemberPersonalInformation, error)
	MatchPartner(memberId string, partnerId string) error
	ListPartner(memberId string) ([]entity.MemberPersonalInformation, error)
}
type partnerFinderUseCase struct {
	partnerRepo    repository.PartnerRepo
	preferenceRepo repository.MemberPreferenceRepo
}

func (m *partnerFinderUseCase) ViewPartner(id string, pageNo int, pageSize int) ([]entity.MemberPersonalInformation, error) {
	pref, intr, err := m.preferenceRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	var ints []string
	for _, i := range intr {
		ints = append(ints, i.InterestId)
	}
	return m.partnerRepo.Find(id, pref.LookingForGender, pref.LookingForDomicile, pref.LookingForStartAge, pref.LookingForEndAge, ints, pageNo, pageSize)
}

func (m *partnerFinderUseCase) MatchPartner(memberId string, partnerId string) error {
	return m.partnerRepo.Create(memberId, partnerId)
}
func (m *partnerFinderUseCase) ListPartner(memberId string) ([]entity.MemberPersonalInformation, error) {
	return m.partnerRepo.FindAll(memberId)
}

func NewPartnerFinderUseCase(partnerRepo repository.PartnerRepo, preferenceRepo repository.MemberPreferenceRepo) PartnerFinderUseCase {
	return &partnerFinderUseCase{partnerRepo: partnerRepo, preferenceRepo: preferenceRepo}
}
