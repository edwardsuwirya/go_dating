package entity

import (
	"time"
)

type MemberUserAccess struct {
	UserName           string     `json:"userName,omitempty"`
	Password           string     `json:"password,omitempty"`
	MemberId           string     `json:"memberId"`
	JoinDate           *time.Time `json:"joinDate,omitempty"`
	VerificationStatus string     `json:"verificationStatus,omitempty"`
}
