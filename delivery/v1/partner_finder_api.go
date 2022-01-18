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

type PartnerFinderApi struct {
	BaseApi
	partnerFinderUseCase usecase.PartnerFinderUseCase
}

func (m *PartnerFinderApi) findPartner() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var findPartner struct {
			httpreq.ByMemberIdReq
			httpreq.ByLimitReq
		}
		jResp, _ := m.ParseRequestQuery(ctx, &findPartner)
		logger.Log.Debug().Msgf("id:%s, page no: %s, page size: %s", findPartner.MemberId, findPartner.PageNo, findPartner.PageSize)

		partnerInfo, err := m.partnerFinderUseCase.ViewPartner(findPartner.MemberId, findPartner.PageNo, findPartner.PageSize)
		if err != nil {
			logger.Log.Err(err).Msg("Find Partner Failed")
			return jResp.SendError(resp.NewErrorMessage(http.StatusInternalServerError, "", "Find Partner Failed"))
		}
		return jResp.SendData(resp.NewResponseMessage("", "Find Partner Success", partnerInfo))
	}
}

func (m *PartnerFinderApi) matchPartner() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(httpreq.MemberPartnerReq)
		jResp, _ := m.ParseRequestBody(ctx, newReq)
		logger.Log.Debug().Msg(newReq.MemberId)
		err := m.partnerFinderUseCase.MatchPartner(newReq.MemberId, newReq.PartnerId)
		if err != nil {
			logger.Log.Err(err).Msg("Match Partner Failed")
			return jResp.SendError(resp.NewErrorMessage(http.StatusInternalServerError, "", "Match Partner Failed"))
		}
		return jResp.SendData(resp.NewResponseMessage("", "Match Partner Success", struct{}{}))
	}
}

func (m *PartnerFinderApi) listMatchPartner() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newReq := new(httpreq.ByMemberIdReq)
		jResp, _ := m.ParseRequestQuery(ctx, newReq)
		logger.Log.Debug().Msg(newReq.MemberId)
		partners, err := m.partnerFinderUseCase.ListPartner(newReq.MemberId)
		if err != nil {
			logger.Log.Err(err).Msg("List Partner Failed")
			return jResp.SendError(resp.NewErrorMessage(http.StatusInternalServerError, "", "List Partner Failed"))
		}
		return jResp.SendData(resp.NewResponseMessage("", "List Partner Success", partners))
	}
}

func NewPartnerFinderApi(rg fiber.Router, partnerFinderUseCase usecase.PartnerFinderUseCase) error {
	if partnerFinderUseCase == nil {
		return errors.New("Empty UseCase")
	}
	partnerFinderApi := &PartnerFinderApi{
		partnerFinderUseCase: partnerFinderUseCase,
	}
	memberGroup := rg.Group("/partner")
	memberGroup.Post("/match", partnerFinderApi.matchPartner())
	memberGroup.Get("/view", partnerFinderApi.findPartner())
	memberGroup.Get("/list", partnerFinderApi.listMatchPartner())
	return nil
}
