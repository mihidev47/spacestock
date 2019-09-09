package util

import (
	"github.com/go-sql-driver/mysql"
	"time"
)

// todayStart returns today epoch that starts from 00:00:00
func TodayStart() int64 {
	return time.Now().Truncate(24 * time.Hour).Unix()
}

// RelativeTodayStart returns today epoch that starts from 00:00:00 + hours of local time difference
func RelativeTodayStartEpoch(diff int8) int64 {
	return TodayStart() + (int64(diff) * 3600)
}

// RelativeEpoch returns now epoch + hours of local time difference
func RelativeEpoch(diff int8) int64 {
	return time.Now().Unix() + (int64(diff) * 3600)
}

// NewNullTime convert time to mysql.NullTime type
func NewNullTime(t time.Time) mysql.NullTime {
	// Time is valid/not null if epoch is more than 0
	valid := t.Unix() > 0
	// Create new null time
	return mysql.NullTime{Time: t, Valid: valid}

}
