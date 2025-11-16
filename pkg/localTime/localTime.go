package localtime

import (
	"time"

	"github.com/rvflash/elapsed"
)

func GetLocalTimeSinceDate(t time.Time) string {
	date := elapsed.LocalTime(t, "ru")
	if date == "еще нет" {
		date = "сейчас"
	}
	return date
}
