package gf

import (
	"math/big"
)

// BasePolynomial calculates the Lagrange base polynomial `l_j` *at position
// 0*, ie `l_j(0)`.
//
// Recall the general Lagrange base polynomial:
// `l_j(x) = Product for m = 0 to k, where m != k [ (x - x_m) / (x_j - x_m) ]
// For x = 0 this simplifies to:
// `l_j(0) = Product for m = 0 to k, where m != k [ x_m / (x_m - x_j) ]
//
// This simplification is sufficient to calculate the value of the polynomial
// at x = 0, which is all we need for retrieving the secret. It has the benefit
// that each evaluate of a base polynomial is a scalar rather than a
// polynomial.
func BasePolynomial(j int, xs []*big.Int, field GF) *big.Int {
	// We'll start with a `1` as it's the identity value of multiplication
	out := big.NewInt(1)
	xj := xs[j]

	for i := 0; i < len(xs); i++ {
		if i == j {
			continue
		}

		term := field.Sub(xs[i], xj)   // x_i - x_j
		term = field.MultInverse(term) // (x_i - x_j)^{-1}
		term = field.Mul(xs[i], term)  // x_i / (x_i - x_j)

		out = field.Mul(out, term)
	}

	return out
}
