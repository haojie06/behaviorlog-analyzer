package data

import "time"

type BlockLog struct {
	Time      time.Time
	Action    string
	Player    string
	Position  string
	PositionX int
	PositionY int
	PositionZ int
	Dimension string
	Target    string
	Remark    string
}
