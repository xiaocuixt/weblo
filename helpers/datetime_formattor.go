package helpers

import (
	"time"
)

// 格式化时间
func DatetimeFormat(time time.Time, format string) string {
	return time.Format(format)
}

func Incr(i int) int {
	return i + 1
}

func Decr(i int) int {
	return i - 1
}
