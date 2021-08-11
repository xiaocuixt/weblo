package helpers

import (
	"time"
)

// 格式化时间
func DatetimeFormat(time time.Time, format string) string {
	return time.Format(format)
}
