package data

import "time"

type LogItem struct {
	Time      time.Time
	LogType   string
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
