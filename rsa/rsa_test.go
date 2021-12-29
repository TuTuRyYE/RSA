package rsa

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	cases := []struct {
		input    int
		expected bool
	}{
		{1, true},
		{4, false},
		{13, true},
		{25, false},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("Is %d a prime number ? %t", test.input, test.expected), func(t *testing.T) {
			assert.Equal(t, test.expected, IsPrimeNumber(test.input))
		})
	}
}

func TestFindTwoPrimeNumbers(t *testing.T) {
	t.Run("numbers are prime", func(t *testing.T) {
		p, q := FindTwoPrimeNumbers()
		assert.Equal(t, true, IsPrimeNumber(p))
		assert.Equal(t, true, IsPrimeNumber(q))
	})
	t.Run("numbers product is higher than 2^24+1", func(t *testing.T) {
		p, q := FindTwoPrimeNumbers()
		assert.Equal(t, true, MIN < p*q)
	})
	t.Run("numbers product is higher than 2^24+1", func(t *testing.T) {
		p, q := FindTwoPrimeNumbers()
		assert.Equal(t, true, p*q < MAX)
	})
}
