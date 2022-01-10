package appresponse

type AppHttpResponse interface {
	SendData(message *ResponseMessage) error
	SendError(errMessage *ErrorMessage) error
}
