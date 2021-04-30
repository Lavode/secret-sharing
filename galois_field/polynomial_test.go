package galois_field

import (
	"math/big"
	"testing"
)

func TestNewPolynomial(t *testing.T) {
	gf := GF{P: big.NewInt(17)}
	poly := NewPolynomial(3, gf)

	if len(poly.Coefficients) != 4 {
		t.Errorf("Expected polynomial of degree 3 to have 4 coefficients; got %d", len(poly.Coefficients))
	}
}

func TestDegree(t *testing.T) {
	gf := GF{P: big.NewInt(17)}
	poly := NewPolynomial(11, gf)

	if poly.Degree() != 11 {
		t.Errorf("Expected polynomial to have degree 11; got %d", poly.Degree())
	}
}

func TestEvaluate(t *testing.T) {
	gf := GF{P: big.NewInt(17)}
	// p(x) = 15 x^2 + 8x + 3
	poly := NewPolynomial(2, gf)
	poly.Coefficients[0] = big.NewInt(3)
	poly.Coefficients[1] = big.NewInt(8)
	poly.Coefficients[2] = big.NewInt(15)

	t.Log("p(x) = 15 x^2 + 8x + 3")
	checks := []struct {
		x int64
		y int64
	}{
		{0, 3},
		{1, 9},
		{2, 11},
		{3, 9},
		{4, 3},
		{5, 10},
		{6, 13},
	}

	for _, check := range checks {
		actual := poly.Evaluate(big.NewInt(check.x))
		if actual.Cmp(big.NewInt(check.y)) != 0 {
			t.Errorf("Expected p(%d) = %d; got %d", check.x, check.y, actual)
		}
	}

	poly = NewPolynomial(0, gf)
	poly.Coefficients[0] = big.NewInt(12)
	t.Log("p(x) = 15 x^2 + 8x + 3")

	actual := poly.Evaluate(big.NewInt(10))
	if actual.Cmp(big.NewInt(12)) != 0 {
		t.Errorf("Expected p(%d) = %d; got %d", 10, 12, actual)
	}
}

func TestString(t *testing.T) {
	gf := GF{P: big.NewInt(17)}
	// p(x) = 15 x^2 + 8x + 3
	poly := NewPolynomial(2, gf)
	poly.Coefficients[0] = big.NewInt(3)
	poly.Coefficients[1] = big.NewInt(8)
	poly.Coefficients[2] = big.NewInt(15)

	expected := "p(x) = 15 x^2 + 8 x^1 + 3"
	if poly.String() != expected {
		t.Errorf("Expected string representation '%s'; Got '%s'", expected, poly.String())
	}

	poly = NewPolynomial(0, gf)
	poly.Coefficients[0] = big.NewInt(12)
	expected = "p(x) = 12"
	if poly.String() != expected {
		t.Errorf("Expected string representation '%s'; Got '%s'", expected, poly.String())
	}
}
