package utils

import (
	"knowledge-api/internal/usecases"
	"log"

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

	_, err := c.AddFunc("*/5 * * * *", func() {
		log.Println("*** start schedule stats ***")
		usecases.StatScheduleUsecase()
	})
	if err != nil {
		log.Println(err)
	}

	c.Start()
}
