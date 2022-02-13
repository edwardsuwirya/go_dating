package manager

import "github.com/edwardsuwirya/go_dating/repository"

type RepositoryManager interface {
	MemberInfoRepo() repository.MemberInfoRepo
	MemberAccessRepo() repository.MemberAccessRepo
	MemberPreferenceRepo() repository.MemberPreferenceRepo
	PartnerRepo() repository.PartnerRepo
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) MemberInfoRepo() repository.MemberInfoRepo {
	return repository.NewMemberPersonalInfoRepo(r.infraManager.GetSqlConn())
}

func (r *repositoryManager) MemberAccessRepo() repository.MemberAccessRepo {
	return repository.NewMemberAccessRepo(r.infraManager.GetSqlConn())
}
func (r *repositoryManager) MemberPreferenceRepo() repository.MemberPreferenceRepo {
	return repository.NewMemberPreferenceRepo(r.infraManager.GetSqlConn())
}
func (r *repositoryManager) PartnerRepo() repository.PartnerRepo {
	return repository.NewPartnerRepo(r.infraManager.GetSqlConn())
}
func NewRepoManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{infraManager: infraManager}
}
