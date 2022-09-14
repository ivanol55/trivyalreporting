package datetimeString

import (
	"strconv";
	"time";
)

func GetCurrentDatetimeString() string {
	var date string = time.Now().Format("2006-01-02")
	var hour, _, _ = time.Now().Clock()
	var hourString = strconv.Itoa(hour)
	var _, minute, _ = time.Now().Clock()
	var minuteString = strconv.Itoa(minute)
	var datetime string = date + "-" + hourString + "-" + minuteString
	return datetime
}