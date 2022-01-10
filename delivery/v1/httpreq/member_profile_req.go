package httpreq

import (
	"fmt"
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/edwardsuwirya/go_dating/util/date"
)

type MemberProfileReq struct {
	MemberId        string
	FirstName       string
	LastName        string
	Bod             string
	Gender          string
	MaritalStatus   string
	Occupation      string
	SelfDescription string
	Instagram       string
	Twitter         string
	MobilePhone     string
	Address         string
	City            string
	PostalCode      string
	Email           string
}

func (p *MemberProfileReq) String() string {
	return fmt.Sprintf("MemberProfileReq => Id: %s, Email: %s", p.MemberId, p.Email)
}

func (p *MemberProfileReq) ToMember() *entity.Member {
	return &entity.Member{
		MemberPersonalInformation: entity.MemberPersonalInformation{
			MemberId:        p.MemberId,
			FirstName:       p.FirstName,
			LastName:        p.LastName,
			Bod:             date.StringToDate(p.Bod),
			Gender:          p.Gender,
			MaritalStatus:   p.MaritalStatus,
			Occupation:      p.Occupation,
			SelfDescription: p.SelfDescription,
		},
		MemberAddressInformation: entity.MemberAddressInformation{
			Address:    p.Address,
			City:       p.City,
			PostalCode: p.PostalCode,
		},
		MemberContactInformation: entity.MemberContactInformation{
			MobilePhoneNumber: p.MobilePhone,
			InstagramId:       p.Instagram,
			TwitterId:         p.Twitter,
			Email:             p.Email,
		},
	}
}
