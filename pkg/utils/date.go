package utils

import "time"

const d = "2006-01-02"
const dt = "2006-01-02 15:04:05"

func NowTime() time.Time {
	return time.Now()
}

// 当期日期时间
func NowDT() string {
	return time.Now().Format(dt)
}

// 当前日期
func NowD() string {
	return time.Now().Format(d)
}

// 转换为日期时间
func FormatDT(time time.Time) string {
	return time.Format(dt)
}

// 转换为日期
func FormatD(time time.Time) string {
	return time.Format(d)
}
