package utils

import (
	"fmt"
	"time"
)

func GetThaiTime() time.Time {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("ERROR: GetThaiTime")
		return time.Now()
	}
	now := time.Now().In(loc)
	return now
}
