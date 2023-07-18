package models

import "time"

type Stat struct {
	Users     int64     `json:"users"`
	Category  int64     `json:"category"`
	Article   int64     `json:"article"`
	CreatedAt time.Time `json:"created_at"`
}
