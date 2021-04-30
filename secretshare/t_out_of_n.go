package secretshare

import (
	"fmt"
	"github.com/lavode/secret-sharing/gf"
	"math/big"
)

// Struct representing a single party's share.
type Share struct {
	Id    int
	Value *big.Int
}

// t-out-of-n secret sharing using polynomials over a finite field GF(p).
//
// It is required that:
// - 1 < t < n
// - t, n are elements of GF(p)
// - The secret is an element of GF(p)
//
// Returns a slice containing the shares and the polynomial used to calculate
// the shares.
// An error is returned if any of the requirements are violated.
func TOutOfN(secret *big.Int, t int, n int, field gf.GF) ([]Share, gf.Polynomial, error) {
	var pol gf.Polynomial
	shares := make([]Share, n)

	if t <= 1 || t >= n || !field.IsGroupElement(big.NewInt(int64(t))) {
		return shares, pol, fmt.Errorf("Invalid value for t")
	}

	if !field.IsGroupElement(big.NewInt(int64(n))) {
		return shares, pol, fmt.Errorf("Invalid value for n")
	}

	if !field.IsGroupElement(secret) {
		return shares, pol, fmt.Errorf("Invalid value for secret")
	}

	pol, err := field.RandomPolynomial(t - 1)
	if err != nil {
		return shares, pol, err
	}

	// We'll use the secret as the first coefficient, so p(0) = secret
	pol.Coefficients[0] = secret

	for i := 0; i < n; i += 1 {
		// Share of participant `i` will be p(i)
		x := i + 1
		result, err := pol.Evaluate(big.NewInt(int64(x)))
		if err != nil {
			return shares, pol, err
		}

		shares[i] = Share{Id: x, Value: result}
	}

	return shares, pol, nil
}
