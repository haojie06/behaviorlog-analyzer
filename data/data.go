package data

import "time"

type BlockLog struct {
	Id        int64
	Time      time.Time
	Action    string
	Player    string
	Position  string
	PositionX int64
	PositionY int64
	PositionZ int64
	Dimension string
	Target    string
	Remark    string
}
