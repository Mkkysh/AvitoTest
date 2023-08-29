package models

import (
	"time"
)

type Log struct {
	Id        int
	IdUser    int `json:"id_user"`
	IdSegment int `json:"id_segment"`
	Operation string
	Timestamp time.Time
}
