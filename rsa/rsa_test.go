package rsa

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	cases := []struct {
		input    int64
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
	t.Run("numbers product is lower than 2^32", func(t *testing.T) {
		p, q := FindTwoPrimeNumbers()
		assert.Equal(t, true, p*q < MAX)
	})
}

func TestFindNextPrimeNumber(t *testing.T) {
	expected := int64(7)
	assert.Equal(t, expected, FindNextPrimeNumber(6))
	assert.Equal(t, expected, FindNextPrimeNumber(7))
}

func TestFindCoPrimeNumber(t *testing.T) {
	t.Run("return a prime number", func(t *testing.T) {
		c := FindCoPrimeNumber(15)
		assert.Equal(t, true, IsPrimeNumber(c))
	})
	t.Run("return a prime number superior of 1", func(t *testing.T) {
		c := FindCoPrimeNumber(15)
		assert.Equal(t, true, 1 < c)
	})
	t.Run("return a prime number inferior of upper bound", func(t *testing.T) {
		var upperBound int64 = 3
		c := FindCoPrimeNumber(upperBound)
		assert.Equal(t, true, c < upperBound)
	})
}

func TestFindD(t *testing.T) {
	t.Run("returns 4", func(t *testing.T) {
		var c, n int64 = 3, 11
		d := FindD(c, n)
		expected := int64(4)
		assert.Equal(t, expected, d)
	})
}

func TestNewKey(t *testing.T) {
	key := NewKey()
	assert.Equal(t, true, key.c*key.d%key.n == 1)
}

func TestKey_Encrypt(t *testing.T) {
	key := NewKeyMock()
	message := "Hello world!"
	//expected := []rune{4744556, 7106336, 7827314, 7103521}
	expected := []uint32{352431401, 2267192425, 538638606, 1131048795}
	assert.Equal(t, expected, key.Encrypt(message))
}

func NewKeyMock() Key {
	return Key{N: 3100069681, n: 3099958000, c: 66797, d: 1336940133}
}

func TestKey_Encrypt3Bytes(t *testing.T) {
	t.Run("from [3]byte to rune", func(t *testing.T) {
		key := Key{2, 1, 1, 1}
		var message = [3]byte{0, 0, 1}
		assert.Equal(t, uint32(1), key.Encrypt3Bytes(message))
	})
	t.Run("from [3]byte to rune", func(t *testing.T) {
		key := Key{2, 1, 1, 1}
		var message = [3]byte{0, 0, 3}
		assert.Equal(t, uint32(1), key.Encrypt3Bytes(message))
	})
}

func TestBytesToUInt32(t *testing.T) {
	cases := []struct {
		input    [3]byte
		expected uint32
	}{
		{[3]byte{0, 0, 0}, 0},
		{[3]byte{0, 0, 255}, 255},
		{[3]byte{0, 1, 0}, 256},
		{[3]byte{1, 0, 0}, 65536},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%v is equal to %d", test.input, test.expected), func(t *testing.T) {
			assert.Equal(t, test.expected, bytesToUInt32(test.input[:]))
		})
	}
}

func TestKey_Decrypt(t *testing.T) {
	key := NewKeyMock()
	message := "Hello world!"
	//expected := []rune{4744556, 7106336, 7827314, 7103521}
	input := []uint32{352431401, 2267192425, 538638606, 1131048795}
	assert.Equal(t, message, key.Decrypt(input))
}

func TestUInt32ToBytes(t *testing.T) {
	cases := []struct {
		input    uint32
		expected [3]byte
	}{
		{0, [3]byte{0, 0, 0}},
		{256, [3]byte{0, 1, 0}},
		{257, [3]byte{0, 1, 1}},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d is equal to %v", test.input, test.expected), func(t *testing.T) {
			assert.Equal(t, test.expected, uInt32toBytes(test.input))
		})
	}
}
