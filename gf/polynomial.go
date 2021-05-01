package gf

import (
	"fmt"
	"math/big"
	"strings"
)

// Polynomial over a finite field
type Polynomial struct {
	// Field the polynomial is in
	Field GF
	// Coefficients of the polynomial, ordered from the lowest degree to
	// the highest
	Coefficients []*big.Int
}

// NewPolynomial initializes a new polynomial of given degree in the given
// field.
func NewPolynomial(degree int, field GF) (Polynomial, error) {
	poly := Polynomial{Field: field}

	if degree < 0 {
		return poly, fmt.Errorf("Degree must be positive")
	}

	poly.Coefficients = make([]*big.Int, degree+1)

	return poly, nil
}

// Degree returns the degree of the polynomial
func (pol *Polynomial) Degree() int {
	return len(pol.Coefficients) - 1
}

// Evaluate the polynomial at a given point
//
// Returns an error if the provided value is not a valid group element.
func (pol *Polynomial) Evaluate(x *big.Int) (*big.Int, error) {
	var result = &big.Int{}

	if !pol.Field.IsGroupElement(x) {
		return result, fmt.Errorf("%d is not a valid group element", x)
	}

	// Coefficient a_i at index i corresponds to a_i * x^i
	for exp, coef := range pol.Coefficients {
		// We'll utilize modular exponentiation for each term of the
		// sum, to prevent having potentially huge intermediary values
		term := pol.Field.Exp(x, big.NewInt(int64(exp))) // x^i
		term = pol.Field.Mul(term, coef)                 // a_i * x^i

		result = pol.Field.Add(result, term) // a_0 + a_1 * x + ... + a_n * x^n
	}

	return result, nil
}

// Return a string representation of this polynomial for printing purposes.
func (pol *Polynomial) String() string {
	var b strings.Builder

	b.WriteString("p(x) = ")
	for i := pol.Degree(); i >= 0; i-- {
		if i == 0 {
			fmt.Fprintf(&b, "%d", pol.Coefficients[i])
		} else {
			fmt.Fprintf(&b, "%d x^%d + ", pol.Coefficients[i], i)
		}
	}

	return b.String()
}
