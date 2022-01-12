package usecase

import (
	"github.com/edwardsuwirya/go_dating/delivery/v1/httpreq"
	"github.com/edwardsuwirya/go_dating/repository"
)

type AuthenticationUseCase interface {
	Login(loginReq *httpreq.LoginReq) (bool, error)
}
type authenticationUseCase struct {
	repo repository.MemberAccessRepo
}

func (m *authenticationUseCase) Login(loginReq *httpreq.LoginReq) (bool, error) {
	return m.repo.FindByUserNameAndPasswordAndVerified(loginReq.UserName, loginReq.Password)
}

func NewAuthenticationUseCase(repo repository.MemberAccessRepo) AuthenticationUseCase {
	return &authenticationUseCase{repo: repo}
}
