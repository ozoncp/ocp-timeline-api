package models

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

type Timeline struct {
	Id     uint64
	UserId uint64
	Type   uint64
	From   timestamp.Timestamp
	To     timestamp.Timestamp
}

func (t *Timeline) String() string {

	fromStr := time.Time(t.From.AsTime()).Format(time.RFC3339)
	toStr := time.Time(t.To.AsTime()).Format(time.RFC3339)

	return fmt.Sprintf("Id: %v; UserId: %v, Type: %v; From: %v; To: %v", t.Id, t.UserId, t.Type, fromStr, toStr)
}
