package domain

import "math/big"

func CalculateFactorial(number uint64) *big.Int {
	factorial := big.NewInt(1)

	for i := uint64(1); i <= number; i++ {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}

	return factorial
}
