package utils

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func ConvertTimeInTimeStamp(time time.Time) timestamp.Timestamp {
	return timestamp.Timestamp{Seconds: time.Unix()}
}
