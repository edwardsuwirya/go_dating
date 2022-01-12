package httpreq

import (
	"fmt"
	"github.com/edwardsuwirya/go_dating/entity"
	"github.com/google/uuid"
	"time"
)

type MemberRegistrationReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *MemberRegistrationReq) String() string {
	return fmt.Sprintf("MemberRegistrationReq => Email: %s, Password: %s", p.Email, p.Password)
}

func (p *MemberRegistrationReq) ToMemberUserAccessForRegistration() *entity.MemberUserAccess {
	joinDate := time.Now().Local()
	return &entity.MemberUserAccess{
		UserName:           p.Email,
		Password:           p.Password,
		MemberId:           uuid.New().String(),
		JoinDate:           &joinDate,
		VerificationStatus: "N",
	}
}
