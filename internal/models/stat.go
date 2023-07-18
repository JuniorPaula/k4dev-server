package models

import "time"

type Stat struct {
	Users     int       `json:"users"`
	Category  int       `json:"category"`
	Article   int       `json:"article"`
	CreatedAt time.Time `json:"created_at"`
}
