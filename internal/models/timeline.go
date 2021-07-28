package models

import (
	"fmt"
	"time"
)

type Timestamp int64

type Timeline struct {
	Id     uint64
	UserId uint64
	Type   uint64
	From   Timestamp
	To     Timestamp
}

func (t *Timeline) String() string {

	fromStr := time.Unix(int64(t.From), 0).Format(time.RFC3339)
	toStr := time.Unix(int64(t.To), 0).Format(time.RFC3339)

	return fmt.Sprintf("Id: %v; UserId: %v, Type: %v; From: %v; To: %v", t.Id, t.UserId, t.Type, fromStr, toStr)
}
