package galois_field

import (
	"crypto/rand"
	"math/big"
)

// Implementation of a finite field *of prime order*. As such only a subset of
// finite fields is supported, corresponding to the ring of integers modulo p.
type GF struct {
	P *big.Int
}

func (gf *GF) Add(a *big.Int, b *big.Int) *big.Int {
	var sum = &big.Int{}
	sum.Add(a, b)
	sum.Mod(sum, gf.P)

	return sum
}

func (gf *GF) Mul(a *big.Int, b *big.Int) *big.Int {
	var prod = &big.Int{}
	prod.Mul(a, b)
	prod.Mod(prod, gf.P)

	return prod
}

func (gf *GF) Rand() (*big.Int, error) {
	return rand.Int(rand.Reader, gf.P)
}
