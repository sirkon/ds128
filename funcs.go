package ds128

import (
	"math/bits"
)

// Add emulates 128bit addition
func Add(lo1, hi1, lo2, hi2 uint64) (lo uint64, hi uint64) {
	lo = lo1 + lo2
	hi = hi1 + hi2
	if lo < lo1 {
		hi++
	}
	return
}

// Mul64 emulates multiplication of 128 bit to 64 bit unsigned integer
func Mul64(lo, hi, v uint64) (resLo uint64, resHi uint64) {
	addHi, resLo := bits.Mul64(lo, v)
	_, resHi = bits.Mul64(hi, v)
	return resLo, resHi + addHi
}

// Mul emulates multiplication of two 128-bit unsigned integers in \mathbb{Z}_{2^128}
func Mul(lo1, hi1, lo2, hi2 uint64) (lo uint64, hi uint64) {
	loLo, hiLo := Mul64(lo2, hi2, lo1)
	loHi, _ := Mul64(lo2, hi2, hi1)
	return loLo, loHi + hiLo
}

// Negate negates 128bit integer
func Negate(lo, hi uint64) (uint64, uint64) {
	return Add(^lo, ^hi, 1, 0)
}

// Cmp checks if first number (lo1, hi1) is less than the second (lo2, hi2)
func Cmp(lo1, hi1, lo2, hi2 uint64) bool {
	if hi1 >= hi2 {
		return false
	}
	if hi1 < hi2 {
		return true
	}
	return lo1 < lo2
}
