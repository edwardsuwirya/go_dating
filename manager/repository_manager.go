package manager

import "github.com/edwardsuwirya/go_dating/repository"

type RepositoryManager interface {
	MemberInfoRepo() repository.MemberInfoRepo
	MemberAccessRepo() repository.MemberAccessRepo
	MemberPreferenceRepo() repository.MemberPreferenceRepo
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) MemberInfoRepo() repository.MemberInfoRepo {
	return repository.NewMemberPersonalInfoRepo(r.infraManager.SqlDb())
}

func (r *repositoryManager) MemberAccessRepo() repository.MemberAccessRepo {
	return repository.NewMemberAccessRepo(r.infraManager.SqlDb())
}
func (r *repositoryManager) MemberPreferenceRepo() repository.MemberPreferenceRepo {
	return repository.NewMemberPreferenceRepo(r.infraManager.SqlDb())
}

func NewRepoManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{infraManager: infraManager}
}
