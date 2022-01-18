package manager

import "github.com/edwardsuwirya/go_dating/usecase"

type UseCaseManager interface {
	MemberRegistrationUseCase() usecase.MemberRegistrationUseCase
	MemberProfileUseCase() usecase.MemberProfileUseCase
	MemberPreferenceUseCase() usecase.MemberPreferenceUseCase
	AuthUseCase() usecase.AuthenticationUseCase
	PartnerFinderUseCase() usecase.PartnerFinderUseCase
}

type useCaseManager struct {
	repo RepositoryManager
}

func (uc *useCaseManager) MemberRegistrationUseCase() usecase.MemberRegistrationUseCase {
	return usecase.NewMemberRegistrationUseCase(uc.repo.MemberAccessRepo())
}
func (uc *useCaseManager) MemberProfileUseCase() usecase.MemberProfileUseCase {
	return usecase.NewMemberProfileUseCase(uc.repo.MemberInfoRepo())
}
func (uc *useCaseManager) MemberPreferenceUseCase() usecase.MemberPreferenceUseCase {
	return usecase.NewMemberPreferenceUseCase(uc.repo.MemberPreferenceRepo())
}
func (uc *useCaseManager) AuthUseCase() usecase.AuthenticationUseCase {
	return usecase.NewAuthenticationUseCase(uc.repo.MemberAccessRepo())
}
func (uc *useCaseManager) PartnerFinderUseCase() usecase.PartnerFinderUseCase {
	return usecase.NewPartnerFinderUseCase(uc.repo.PartnerRepo(), uc.repo.MemberPreferenceRepo())
}

func NewUseCaseManger(manager RepositoryManager) UseCaseManager {
	return &useCaseManager{repo: manager}
}
