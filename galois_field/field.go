// Package galois_field implements operations over finite fields *of prime
// order*. As such only a subset of fields is supported, each corresponding to
// the ring of integers modulo p.
package galois_field

import (
	"crypto/rand"
	"math/big"
)

// Finite field of prime order
type GF struct {
	P *big.Int
}

// Addition in the finite field `gf`.
func (gf *GF) Add(a *big.Int, b *big.Int) *big.Int {
	var sum = &big.Int{}
	sum.Add(a, b)
	sum.Mod(sum, gf.P)

	return sum
}

// Multiplication in the finite field `gf`.
func (gf *GF) Mul(a *big.Int, b *big.Int) *big.Int {
	var prod = &big.Int{}
	prod.Mul(a, b)
	prod.Mod(prod, gf.P)

	return prod
}

// Get a random rember of the finite field `gf`.
func (gf *GF) Rand() (*big.Int, error) {
	return rand.Int(rand.Reader, gf.P)
}

func (gf *GF) RandomPolynomial(degree int) (Polynomial, error) {
	poly := NewPolynomial(degree, *gf)

	for i := 0; i < degree+1; i += 1 {
		rnd, err := gf.Rand()
		if err != nil {
			return poly, err
		}

		poly.Coefficients[i] = rnd
	}

	return poly, nil
}
