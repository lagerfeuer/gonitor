package util

import (
	"strconv"
	"strings"
	"time"
)

func UptimeToHumanReadable(uptimeSecs uint64) string {
	ts := time.Unix(int64(uptimeSecs), 0)
	builder := strings.Builder{}
	if ts.Day() > 0 {
		builder.WriteString(strconv.Itoa(ts.Day()) + " days ")
	}
	if builder.Len() > 0 || ts.Hour() > 0 {
		builder.WriteString(strconv.Itoa(ts.Hour()) + " hours ")
	}
	if builder.Len() > 0 || ts.Minute() > 0 {
		builder.WriteString(strconv.Itoa(ts.Minute()) + " minutes ")
	}
	if ts.Day() > 0 || ts.Second() > 0 {
		builder.WriteString(strconv.Itoa(ts.Second()) + " seconds ")
	}
	return builder.String()
}
