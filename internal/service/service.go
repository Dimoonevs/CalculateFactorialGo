package service

import (
	"math/big"
	"sync"

	"github.com/Dimoonevs/calculate/factorial/internal/domain"
	"github.com/Dimoonevs/calculate/factorial/internal/models"
)

type Service struct{}

type ServiceInterface interface {
	CalculateFactorial(*models.JsonPayload) (big.Int, big.Int)
}

func (s *Service) CalculateFactorial(data *models.JsonPayload) (big.Int, big.Int) {
	var wg sync.WaitGroup

	var factorialA, factorialB *big.Int

	wg.Add(2)

	go func() {
		defer wg.Done()
		factorialA = domain.CalculateFactorial(*data.A)
	}()
	go func() {
		defer wg.Done()
		factorialB = domain.CalculateFactorial(*data.B)
	}()
	wg.Wait()

	return *factorialA, *factorialB
}

func NewService() ServiceInterface {
	return &Service{}
}
