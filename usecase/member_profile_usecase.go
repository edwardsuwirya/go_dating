package usecase

import (
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/repository"
)

type MemberProfileUseCase interface {
	UpdateProfile(member *entity.Member) (*entity.Member, error)
	GetProfile(id string) (*entity.Member, error)
}
type memberProfileUseCase struct {
	infoRepo repository.MemberInfoRepo
}

func (m *memberProfileUseCase) UpdateProfile(member *entity.Member) (*entity.Member, error) {
	return m.infoRepo.Create(member)
}

func (m *memberProfileUseCase) GetProfile(id string) (*entity.Member, error) {
	return m.infoRepo.FindById(id)
}
func NewMemberProfileUseCase(repo repository.MemberInfoRepo) MemberProfileUseCase {
	return &memberProfileUseCase{infoRepo: repo}
}
