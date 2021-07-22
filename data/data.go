package data

import "time"

type BlockLog struct {
	Id        int64
	Time      time.Time
	Action    string
	Player    string
	Position  string
	PositionX int32
	PositionY int32
	PositionZ int32
	Dimension string
	Target    string
	Remark    string
}
