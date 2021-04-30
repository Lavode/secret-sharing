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
	sum.Add(a, b)      // a + b
	sum.Mod(sum, gf.P) // a + b mod p

	return sum
}

// Subtraction in the finite field `gf`.
func (gf *GF) Sub(a *big.Int, b *big.Int) *big.Int {
	var diff = &big.Int{}
	diff.Sub(a, b)       // a - b
	diff.Mod(diff, gf.P) // a - b mod p

	return diff
}

// Multiplication in the finite field `gf`.
func (gf *GF) Mul(a *big.Int, b *big.Int) *big.Int {
	var prod = &big.Int{}
	prod.Mul(a, b)       // a * b
	prod.Mod(prod, gf.P) // a * b mod p

	return prod
}

// Division in the finite field `gf`.
func (gf *GF) Div(a *big.Int, b *big.Int) *big.Int {
	var quot = &big.Int{}
	var inv = &big.Int{}

	inv.Mod(b, gf.P)     // b^{-1}
	quot.Mul(a, inv)     // a * b^{-1}
	quot.Mod(quot, gf.P) // a * b^{-1} mod p

	return quot
}

// Exponentiation in the finite field `gf`.
func (gf *GF) Exp(b *big.Int, e *big.Int) *big.Int {
	var pow = &big.Int{}
	pow.Exp(b, e, gf.P) // b^e mod p

	return pow
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
