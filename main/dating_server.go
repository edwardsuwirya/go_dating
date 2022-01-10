package main

import (
	v1 "github.com/edwardsuwirya/go_dating/delivery/v1"
	"github.com/edwardsuwirya/go_dating/manager"
	"github.com/edwardsuwirya/go_dating/util/logger"
	"github.com/gofiber/fiber/v2"
)

type DatingServer interface {
	Run()
}

type datingServer struct {
	infra      manager.InfraManager
	useCase    manager.UseCaseManager
	router     *fiber.App
	apiVersion string
}

func (d *datingServer) handlers() {
	apiMainGroup := d.infra.Config().GetString("datingapp.api.group")
	api := d.router.Group(apiMainGroup)
	apiVersionGroup := api.Group(d.apiVersion)
	switch d.apiVersion {
	case "/v1":
		d.v1(apiVersionGroup)
	default:
		logger.Log.Fatal().Msg("Unknown API version")
	}

}
func (d *datingServer) v1(rg fiber.Router) {
	err := v1.NewMemberApi(rg, d.useCase.MemberRegistrationUseCase(), d.useCase.MemberProfileUseCase(), d.useCase.MemberPreferenceUseCase())
	if err != nil {
		logger.Log.Fatal().Msg("Member Registration failed to start")
	}

}
func (d *datingServer) Run() {
	apiUrl := d.infra.Config().GetString("datingapp.api.url")
	d.handlers()
	logger.Log.Info().Msg("Server is running")
	err := d.router.Listen(apiUrl)
	if err != nil {
		logger.Log.Fatal().Msg("Server is failed")
	}
}
func NewDatingServer() DatingServer {
	infra := manager.NewInfra()
	repo := manager.NewRepoManager(infra)
	useCase := manager.NewUseCaseManger(repo)
	fiberRouter := fiber.New(fiber.Config{
		AppName:       infra.Config().GetString("datingapp.name"),
		StrictRouting: true,
		CaseSensitive: true,
	})
	return &datingServer{
		infra:      infra,
		useCase:    useCase,
		router:     fiberRouter,
		apiVersion: infra.Config().GetString("datingapp.api.version"),
	}
}
