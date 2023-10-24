package service

type DaysCount interface {
	CalculateDaysCount() (int16, error)
}

type Service struct {
	DaysCount
}

func NewService() *Service {
	return &Service{
		DaysCount: NewDaysCountService(),
	}
}
