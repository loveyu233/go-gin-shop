package utils

import "time"

type TimeUtils struct {
}

func (TimeUtils) StringTimeToUnix(timeStr string) (int64, error) {
	location, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return -1, err
	}
	return location.Unix(), nil
}

func (t TimeUtils) TimeIsOutTime(timeStartStr, timeEndStr string) bool {
	timeStartStrUnix, err := t.StringTimeToUnix(timeStartStr)
	if err != nil {
		return false
	}
	timeEndStrUnix, err2 := t.StringTimeToUnix(timeEndStr)
	if err2 != nil {
		return false
	}
	timeNowUnix := time.Now().Unix()
	if timeNowUnix >= timeStartStrUnix && timeNowUnix <= timeEndStrUnix {
		return true
	}
	return false
}
