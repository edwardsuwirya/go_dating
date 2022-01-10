package httpreq

import (
	"fmt"
	"github.com/edwardsuwirya/go_dating/entity"
	"time"
)

type MemberRegistrationReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *MemberRegistrationReq) String() string {
	return fmt.Sprintf("MemberRegistrationReq => Email: %s, Password: %s", p.Email, p.Password)
}

func (p *MemberRegistrationReq) ToMemberUserAccess() *entity.MemberUserAccess {
	return &entity.MemberUserAccess{
		UserName:           p.Email,
		Password:           p.Password,
		JoinDate:           time.Now().Local(),
		VerificationStatus: "N",
	}
}
