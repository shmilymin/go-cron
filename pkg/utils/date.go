package utils

import "time"

const d = "2006-01-02"
const dt = "2006-01-02 15:04:05"

func Now() string {
	return time.Now().Format(dt)
}

func Today() string {
	return time.Now().Format(d)
}
