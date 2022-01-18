package httpreq

type ByLimitReq struct {
	PageNo   int `query:"pageNo"`
	PageSize int `query:"pageSize"`
}
