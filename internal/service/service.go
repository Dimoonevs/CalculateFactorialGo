package service

import (
	"sync"

	"github.com/Dimoonevs/calculate/factorial/internal/models"
)

type Service struct{}

type ServiceInterface interface {
	CalculateFactorial(*models.JsonPayload) (uint64, uint64)
}

func (s *Service) CalculateFactorial(data *models.JsonPayload) (uint64, uint64) {
	var wg sync.WaitGroup

	var factorialA, factorialB uint64

	wg.Add(2)

	go calculateFactorial(*data.A, &factorialA, &wg)
	go calculateFactorial(*data.B, &factorialB, &wg)
	wg.Wait()

	return factorialA, factorialB
}

func NewService() ServiceInterface {
	return &Service{}
}

func calculateFactorial(number uint64, result *uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	factorial := uint64(1)

	for i := uint64(1); i <= number; i++ {
		factorial *= i
	}

	*result = factorial
}
