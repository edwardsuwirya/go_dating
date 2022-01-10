package httpreq

import (
	"fmt"
	"github.com/edwardsuwirya/go_dating/entity"
)

type MemberPreferenceReq struct {
	MemberId         string
	GenderInterest   string
	DomicileInterest string
	StartAgeInterest int
	EndAgeInterest   int
}

func (p *MemberPreferenceReq) String() string {
	return fmt.Sprintf("MemberPreferenceReq => Id: %s, Email: %s", p.MemberId, p.GenderInterest)
}
func (p *MemberPreferenceReq) ToMemberPreference() *entity.MemberPreferences {
	return &entity.MemberPreferences{
		MemberId:           p.MemberId,
		LookingForGender:   p.GenderInterest,
		LookingForDomicile: p.DomicileInterest,
		LookingForStartAge: p.StartAgeInterest,
		LookingForEndAge:   p.EndAgeInterest,
	}
}
