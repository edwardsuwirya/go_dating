package httpreq

import "fmt"

type MemberPartnerReq struct {
	MemberId  string `json:"memberId"`
	PartnerId string `json:"partnerId"`
}

func (p *MemberPartnerReq) String() string {
	return fmt.Sprintf("MemberPartnerReq => Member ID: %s, Partner ID: %s", p.MemberId, p.PartnerId)
}
