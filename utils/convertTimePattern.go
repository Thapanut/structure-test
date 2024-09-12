package utils

import (
	"time"
)

func ConvertTimePattern(t time.Time) string {
	tString := t.Format("2006-01-02T15:04:05.999999+07:00")
	return tString
}
