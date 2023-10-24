package service

import "time"

type DaysCountService struct {
}

func NewDaysCountService() *DaysCountService {
	return &DaysCountService{}
}

func (s *DaysCountService) CalculateDaysCount() (int16, error) {
	initTime := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	daysLeft := time.Until(initTime)

	return int16(daysLeft.Hours()) / 24, nil
}
