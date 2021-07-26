package models

type Timestamp int64

type Timeline struct {
	Id     uint64
	UserId uint64
	Type   uint64
	From   Timestamp
	To     Timestamp
}
