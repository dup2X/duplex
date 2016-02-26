package dtime

import (
	"time"
)

func GetMonthStart(ago int) (monthStart time.Time) {
	n := time.Now()
	y, m, _ := n.Date()
	if int(m)-ago < 1 {
		m = time.Month(12 - ago)
		y--
	} else {
		m = time.Month(int(m) - ago)
	}
	monthStart = time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	return
}

func GetDayStart(ago int) (dayStart time.Time) {
	t := time.Now().UnixNano() - int64(3600*24*ago*1000)
	return time.Unix(0, t)
}
