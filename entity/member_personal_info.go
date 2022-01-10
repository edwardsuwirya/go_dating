package entity

import "time"

type MemberPersonalInformation struct {
	MemberId        string
	FirstName       string
	LastName        string
	Bod             time.Time
	Gender          string
	MaritalStatus   string
	Occupation      string
	RecentPhotoPath string
	SelfDescription string
}
