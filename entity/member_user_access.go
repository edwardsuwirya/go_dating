package entity

import (
	"time"
)

type MemberUserAccess struct {
	UserName           string
	Password           string
	MemberId           string
	JoinDate           time.Time
	VerificationStatus string
}
