package v1

import (
	"errors"
	"github.com/edwardsuwirya/go_dating/delivery/v1/httpreq"
	resp "github.com/edwardsuwirya/go_dating/delivery/v1/httpresp"
	"github.com/edwardsuwirya/go_dating/usecase"
	"github.com/edwardsuwirya/go_dating/util/logger"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthApi struct {
	BaseApi
	authUseCase usecase.AuthenticationUseCase
}

func (m *AuthApi) userLogin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(httpreq.LoginReq)
		jResp, _ := m.ParseRequestBody(ctx, newReq)
		logger.Log.Debug().Msg(newReq.String())
		isAuth, err := m.authUseCase.Login(newReq)
		if isAuth == false || err != nil {
			logger.Log.Error().Err(err).Msg("Unauthorized")
			return jResp.SendError(resp.NewErrorMessage(http.StatusUnauthorized, "", "Unauthorized"))
		}
		return jResp.SendData(resp.NewResponseMessage("", "Authentication Success", nil))
	}
}

func NewAuthApi(rg fiber.Router, authUseCase usecase.AuthenticationUseCase) error {
	if authUseCase == nil {
		return errors.New("Empty UseCase")
	}
	authApi := &AuthApi{
		authUseCase: authUseCase,
	}
	memberGroup := rg.Group("/auth")
	memberGroup.Post("", authApi.userLogin())
	return nil
}
