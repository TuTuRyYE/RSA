package rsa

import (
	"math"
)

const MAX = 4294967296 //2^32
const MIN = 16777217   //2^24+1

func IsPrimeNumber(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))+1); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindTwoPrimeNumbers() (int, int) {
	GeneratePrimeNumbers()
	var p = 1
	var q = 1
	for !RespectRules(p, q) {
	}
	return p, q
}

func RespectRules(p int, q int) bool {
	return IsPrimeNumber(p) && IsPrimeNumber(q)
}

func GeneratePrimeNumbers() []int {
	var primeNumbers []int
	for i := 2; i < MAX/2; i++ {
		var isPrime = true
		for _, primeNumber := range primeNumbers {
			if i%primeNumber == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeNumbers = append(primeNumbers, i)
		}

	}
	return primeNumbers
}
