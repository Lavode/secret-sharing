package galois_field

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

// Initialize a new polynomial of given degreen in given field.
func NewPolynomial(degree int, field GF) Polynomial {
	poly := Polynomial{Field: field}
	poly.Coefficients = make([]*big.Int, degree+1)

	return poly
}

// Return the degree of the polynomial
func (pol *Polynomial) Degree() int {
	return len(pol.Coefficients) - 1
}

// Evaluate the polynomial at a given point
func (pol *Polynomial) Evaluate(x *big.Int) *big.Int {
	var result = &big.Int{}

	// Coefficient a_i at index i corresponds to a_i * x^i
	for exp, coef := range pol.Coefficients {
		// We'll utilize modular exponentiation for each term of the
		// sum, to prevent having potentially huge intermediary values
		var term = &big.Int{}
		term.Set(x)                                         // x
		term.Exp(term, big.NewInt(int64(exp)), pol.Field.P) // x^i
		term.Mul(term, coef)                                // a_i * x^i

		result.Add(result, term) // a_0 + a_1 * x + ... + a_n * x^n
	}

	result.Mod(result, pol.Field.P)

	return result
}

func (pol *Polynomial) String() string {
	var b strings.Builder

	b.WriteString("p(x) = ")
	for i := pol.Degree(); i >= 0; i -= 1 {
		if i == 0 {
			fmt.Fprintf(&b, "%d", pol.Coefficients[i])
		} else {
			fmt.Fprintf(&b, "%d x^%d + ", pol.Coefficients[i], i)
		}
	}

	return b.String()
}
