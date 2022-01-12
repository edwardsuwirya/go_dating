package entity

import "time"

type MemberPersonalInformation struct {
	PersonalInformationId string
	MemberId              string
	Name                  string
	Bod                   time.Time
	Gender                string
	RecentPhotoPath       string
	SelfDescription       string
}
