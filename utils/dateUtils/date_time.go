package dateUtils

import (
	"time"
)

const (
	apiDateLayOut = "02-01-2006T15:04:05Z"
	apiDbLayout   = "02-01-2006 15:04:05"
)

func GetNOW() string {
	now := time.Now()

	return now.Format(apiDateLayOut)
}

func GetNowDb() string {
	now := time.Now()

	return now.Format(apiDbLayout)
}
