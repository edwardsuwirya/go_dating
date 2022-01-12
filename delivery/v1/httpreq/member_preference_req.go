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
	Interests        []MemberInterestReq
}

type MemberInterestReq struct {
	InterestId string
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
func (p *MemberPreferenceReq) ToMemberInterest() []entity.MemberInterest {
	var interests []entity.MemberInterest
	for _, i := range p.Interests {
		interests = append(interests, entity.MemberInterest{
			InterestId: i.InterestId,
			MemberId:   p.MemberId,
		})
	}
	return interests
}
