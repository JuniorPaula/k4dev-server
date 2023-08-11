package utils

import (
	"fmt"
	"knowledge-api/internal/usecases"

	"github.com/robfig/cron/v3"
)

type Schedule struct {
	Cron *cron.Cron
}

func NewSchedule() *Schedule {
	return &Schedule{
		Cron: cron.New(),
	}
}

func (s *Schedule) StatsSchedule() {
	c := s.Cron

	_, err := c.AddFunc("*/1 * * * *", func() {
		fmt.Println("start schedule stats")
		usecases.StatScheduleUsecase()
	})
	if err != nil {
		fmt.Println(err)
	}

	c.Start()
}
