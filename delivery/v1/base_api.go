package v1

import (
	resp "github.com/edwardsuwirya/go_dating/delivery/v1/httpresp"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type BaseApi struct {
}

func (b *BaseApi) ParseRequestBody(ctx *fiber.Ctx, newReq interface{}) (resp.AppHttpResponse, error) {
	jResp := resp.NewJsonResponse(ctx)
	if err := ctx.BodyParser(newReq); err != nil {
		return nil, jResp.SendError(resp.NewErrorMessage(http.StatusBadRequest, "", "Request can not be parsed"))
	}
	return jResp, nil
}
func (b *BaseApi) ParseRequestQuery(ctx *fiber.Ctx, newReq interface{}) (resp.AppHttpResponse, error) {
	jResp := resp.NewJsonResponse(ctx)
	if err := ctx.QueryParser(newReq); err != nil {
		return nil, jResp.SendError(resp.NewErrorMessage(http.StatusBadRequest, "", "Request can not be parsed"))
	}
	return jResp, nil
}
