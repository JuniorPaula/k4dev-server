package models

import "time"

type Stat struct {
	Users     int64      `json:"users"`
	Category  int64      `json:"category"`
	Article   int64      `json:"article"`
	CreatedAt time.Timer `json:"created_at"`
}
