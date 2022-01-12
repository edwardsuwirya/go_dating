package httpreq

import (
	"fmt"
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/util/date"
)

type MemberProfileReq struct {
	MemberId        string
	Name            string
	Bod             string
	Gender          string
	SelfDescription string
	Instagram       string
	Twitter         string
	MobilePhone     string
	Address         string
	City            string
	PostalCode      string
}

func (p *MemberProfileReq) String() string {
	return fmt.Sprintf("MemberProfileReq => Id: %s", p.MemberId)
}

func (p *MemberProfileReq) ToMember() *entity.Member {
	return &entity.Member{
		PersonalInfo: entity.MemberPersonalInformation{
			MemberId:        p.MemberId,
			Name:            p.Name,
			Bod:             date.StringToDate(p.Bod),
			Gender:          p.Gender,
			SelfDescription: p.SelfDescription,
		},
		AddressInfo: entity.MemberAddressInformation{
			Address:    p.Address,
			City:       p.City,
			PostalCode: p.PostalCode,
		},
		ContactInfo: entity.MemberContactInformation{
			MobilePhoneNumber: p.MobilePhone,
			InstagramId:       p.Instagram,
			TwitterId:         p.Twitter,
		},
	}
}
