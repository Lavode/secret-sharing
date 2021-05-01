package secretshare

import (
	"fmt"
	"github.com/lavode/secret-sharing/gf"
	"math/big"
)

// Share represents a single party's share of a secret.
type Share struct {
	ID    int
	Value *big.Int
}

// TOutOfN implements t-out-of-n secret sharing using polynomials over a finite
// field GF(p).
//
// It is required that:
// - 1 < t <= n
// - secret, t, n are elements of GF(p)
//
// Returns a slice containing the shares and the polynomial used to calculate
// the shares.
// An error is returned if any of the requirements are violated.
func TOutOfN(secret *big.Int, t int, n int, field gf.GF) ([]Share, gf.Polynomial, error) {
	var pol gf.Polynomial
	shares := make([]Share, n)

	if t <= 1 || t > n || !field.IsGroupElement(big.NewInt(int64(t))) {
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

	for i := 0; i < n; i++ {
		// Share of participant `i` will be p(i)
		x := i + 1
		result, err := pol.Evaluate(big.NewInt(int64(x)))
		if err != nil {
			return shares, pol, err
		}

		shares[i] = Share{ID: x, Value: result}
	}

	return shares, pol, nil
}

// TOutOfNRecover recovers a secret from t out of n shares.
//
// The slice of shares must be *exactly* `t` *unique* shares. If there are any
// more or less, an incorrect value will be reconstructed.
//
// Returns an error if shares are not unique.
func TOutOfNRecover(shares []Share, field gf.GF) (*big.Int, error) {
	var sum = &big.Int{}

	seen := make(map[int]bool)
	xs := make([]*big.Int, len(shares))
	for i, share := range shares {
		if _, ok := seen[share.ID]; ok {
			// Share with given ID already seen
			return sum, fmt.Errorf("Duplicate share with ID %d supplied", share.ID)
		}
		seen[share.ID] = true
		xs[i] = big.NewInt(int64(share.ID))
	}

	for j := 0; j < len(shares); j++ {
		var term = &big.Int{}
		term.Set(shares[j].Value)                   // y_i
		basePoly := gf.BasePolynomial(j, xs, field) // l_j(0)
		term = field.Mul(term, basePoly)            // y_i * l_j(0)
		sum = field.Add(sum, term)
	}

	return sum, nil
}
