// Package gf implements operations over finite fields *of prime order*. As
// such only a subset of fields is supported, each corresponding to the ring of
// integers modulo p.
package gf

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

// Get a random polynomial over the finite field `gf`.
func (gf *GF) RandomPolynomial(degree int) (Polynomial, error) {
	poly, err := NewPolynomial(degree, *gf)

	if err != nil {
		return poly, err
	}

	for i := 0; i < degree+1; i += 1 {
		rnd, err := gf.Rand()
		if err != nil {
			return poly, err
		}

		poly.Coefficients[i] = rnd
	}

	return poly, nil
}

// Check if an element is an element of group `gf`.
func (gf *GF) IsGroupElement(x *big.Int) bool {
	// x < 0
	if x.Cmp(big.NewInt(0)) == -1 {
		return false
	}

	// x >= p
	if x.Cmp(gf.P) != -1 {
		return false
	}

	return true
}
