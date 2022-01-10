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

type MemberApi struct {
	BaseApi
	memberRegistrationUseCase usecase.MemberRegistrationUseCase
	memberProfileUseCase      usecase.MemberProfileUseCase
	memberPreferenceUseCase   usecase.MemberPreferenceUseCase
}

func (m *MemberApi) memberRegistration() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(httpreq.MemberRegistrationReq)
		jResp, _ := m.ParseRequestBody(ctx, newReq)
		logger.Log.Debug().Msg(newReq.String())
		newMember := newReq.ToMemberUserAccess()
		err := m.memberRegistrationUseCase.NewRegistration(newMember)
		if err != nil {
			logger.Log.Err(err).Msg("Registration Failed")
			return jResp.SendError(resp.NewErrorMessage(http.StatusInternalServerError, "", "Registration Failed"))
		}
		return jResp.SendData(resp.NewResponseMessage("", "Member Registration Success", newMember))
	}
}

func (m *MemberApi) memberActivation() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(httpreq.ByMemberIdReq)
		jResp, _ := m.ParseRequestQuery(ctx, newReq)
		logger.Log.Debug().Msg(newReq.MemberId)

		err := m.memberRegistrationUseCase.NewActivation(newReq.MemberId)
		if err != nil {
			logger.Log.Err(err).Msg("Activation Failed")
			return jResp.SendError(resp.NewErrorMessage(http.StatusInternalServerError, "", "Activation Failed"))
		}
		return jResp.SendData(resp.NewResponseMessage("", "Member Activation Success", newReq.MemberId))
	}
}

func (m *MemberApi) getProfile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(httpreq.ByMemberIdReq)
		jResp, _ := m.ParseRequestQuery(ctx, newReq)
		logger.Log.Debug().Msg(newReq.MemberId)
		profile, err := m.memberProfileUseCase.GetProfile(newReq.MemberId)
		if err != nil {
			logger.Log.Err(err).Msg("Get Profile Failed")
			return jResp.SendError(resp.NewErrorMessage(http.StatusInternalServerError, "", "Get Profile Failed"))
		}
		return jResp.SendData(resp.NewResponseMessage("", "Member Profile", profile))
	}
}

func (m *MemberApi) updateProfile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(httpreq.MemberProfileReq)
		jResp, _ := m.ParseRequestBody(ctx, newReq)
		logger.Log.Debug().Msg(newReq.String())
		memberInfo := newReq.ToMember()
		member, err := m.memberProfileUseCase.UpdateProfile(memberInfo)
		if err != nil {
			logger.Log.Err(err).Msg("Update profile Failed")
			return jResp.SendError(resp.NewErrorMessage(http.StatusInternalServerError, "", "Update profile Failed"))
		}
		return jResp.SendData(resp.NewResponseMessage("", "Update profile Success", member))
	}
}

func (m *MemberApi) updatePreference() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(httpreq.MemberPreferenceReq)
		jResp, _ := m.ParseRequestBody(ctx, newReq)
		logger.Log.Debug().Msg(newReq.String())
		memberPref := newReq.ToMemberPreference()
		member, err := m.memberPreferenceUseCase.UpdatePreference(memberPref)
		if err != nil {
			logger.Log.Err(err).Msg("Update profile Failed")
			return jResp.SendError(resp.NewErrorMessage(http.StatusInternalServerError, "", "Update profile Failed"))
		}
		return jResp.SendData(resp.NewResponseMessage("", "Update profile Success", member))
	}
}

func NewMemberApi(rg fiber.Router, memberRegistrationUseCase usecase.MemberRegistrationUseCase, memberProfileUseCase usecase.MemberProfileUseCase, memberPreferenceUseCase usecase.MemberPreferenceUseCase) error {
	if memberRegistrationUseCase == nil || memberPreferenceUseCase == nil || memberProfileUseCase == nil {
		return errors.New("Empty UseCase")
	}
	memberApi := &MemberApi{
		memberRegistrationUseCase: memberRegistrationUseCase,
		memberProfileUseCase:      memberProfileUseCase,
		memberPreferenceUseCase:   memberPreferenceUseCase,
	}
	memberGroup := rg.Group("/member")
	memberGroup.Post("/registration", memberApi.memberRegistration())
	memberGroup.Get("/activation", memberApi.memberActivation())
	memberGroup.Put("/profile", memberApi.updateProfile())
	memberGroup.Get("/profile", memberApi.getProfile())
	memberGroup.Post("/preference", memberApi.updatePreference())
	return nil
}
