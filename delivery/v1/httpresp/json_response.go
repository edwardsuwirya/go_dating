package appresponse

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type JsonResponse struct {
	ctx *fiber.Ctx
}

func (j *JsonResponse) SendData(message *ResponseMessage) error {
	return j.ctx.Status(http.StatusOK).JSON(message)
}

func (j *JsonResponse) SendError(errMessage *ErrorMessage) error {
	return j.ctx.Status(errMessage.HttpCode).JSON(errMessage)
}

func NewJsonResponse(ctx *fiber.Ctx) AppHttpResponse {
	return &JsonResponse{ctx: ctx}
}
