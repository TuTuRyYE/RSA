package rsa

import (
	"math"
	"math/big"
	"math/rand"
)

const MAX = 4294967296 //2^32
const MIN = 16777217   //2^24+1

func IsPrimeNumber(n int64) bool {
	for i := int64(2); i < int64(math.Sqrt(float64(n))+1); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindTwoPrimeNumbers() (int64, int64) {
	var a = rand.Int63n(MAX-MIN) + MIN
	var p = FindNextPrimeNumber(a)
	var b = rand.Int63n((MAX-MIN)/p) + MIN/p
	var q = FindNextPrimeNumber(b)
	if findTwoPrimeNumbersRules(p, q) {
		return p, q
	} else {
		return FindTwoPrimeNumbers()
	}
}

func findTwoPrimeNumbersRules(p int64, q int64) bool {
	return MIN < p*q && p*q < MAX && p > 1 && q > 1
}

func FindNextPrimeNumber(n int64) int64 {
	var primeNumber = n
	for !IsPrimeNumber(primeNumber) {
		primeNumber++
	}
	return primeNumber
}

func FindCoPrimeNumber(n int64) int64 {
	var a = rand.Int63n(n)
	var c = FindNextPrimeNumber(a)
	if 1 < c && c < n {
		return c
	} else {
		return FindCoPrimeNumber(n)
	}
}

func FindD(c int64, n int64) int64 {
	var d int64 = 2
	for c*d%n != 1 {
		d++
	}
	return d
}

type Key struct {
	N, n, c, d int64
}

func (k Key) Encrypt(message string) []uint32 {
	bytes := []byte(message)
	var encryptedMessage []uint32
	for i := 0; i < len(bytes); i += 3 {
		partialBytes := [3]byte{bytes[i], bytes[i+1], bytes[i+2]}
		encryptedMessage = append(encryptedMessage, k.Encrypt3Bytes(partialBytes))
	}
	return encryptedMessage
}

func (k Key) Encrypt3Bytes(message [3]byte) uint32 {
	result := big.NewInt(0)
	a := big.NewInt(int64(bytesToUInt32(message[:])))
	result.Exp(a, big.NewInt(k.c), big.NewInt(k.N))
	return uint32(result.Int64())
}

func (k Key) Decrypt3Bytes(n uint32) [3]byte {
	result := big.NewInt(0)
	a := big.NewInt(int64(n))
	result.Exp(a, big.NewInt(k.d), big.NewInt(k.N))
	return uInt32toBytes(uint32(result.Int64()))
}

func (k Key) Decrypt(input []uint32) string {
	var resultAsByte []byte
	for _, integer := range input {
		for _, b := range k.Decrypt3Bytes(integer) {
			resultAsByte = append(resultAsByte, b)
		}
	}
	return string(resultAsByte)
}

func bytesToUInt32(bytes []byte) uint32 {
	var a uint32
	for k, v := range bytes {
		var b = uint32(math.Pow(256, float64(len(bytes)-1-k)))
		a += uint32(v) * b
	}
	return a
}

func uInt32toBytes(n uint32) [3]byte {
	var result [3]byte
	for k := range result {
		var b = uint32(math.Pow(256, float64(2-k)))
		result[k] = byte(n / b)
		n = n % b
	}
	return result
}

func NewKey() Key {
	p, q := FindTwoPrimeNumbers()
	N := p * q
	n := (p - 1) * (q - 1)
	c := FindCoPrimeNumber(n)
	d := FindD(c, n)
	return Key{N, n, c, d}
}
