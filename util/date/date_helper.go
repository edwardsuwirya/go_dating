package date

import "time"

const (
	layoutISO = "2006-01-02"
)

func StringToDate(stringDate string) time.Time {
	t, _ := time.Parse(layoutISO, stringDate)
	return t
}
