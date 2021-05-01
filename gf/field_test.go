package gf

import (
	"math/big"
	"testing"
)

func TestNewGF(t *testing.T) {
	gf, err := NewGF(big.NewInt(53))
	if err != nil {
		t.Errorf("Error while creating new GF of prime order: %v", err)
	}
	if gf.P.Cmp(big.NewInt(53)) != 0 {
		t.Errorf("Expected order of 53; got %d", gf.P)
	}

	_, err = NewGF(big.NewInt(1024))
	if err == nil {
		t.Errorf("Expected error when creating field of non-prime order; got none")
	}

}

func TestAdd(t *testing.T) {
	gf, err := NewGF(big.NewInt(17))
	if err != nil {
		t.Fatalf("Error generating finite field: %v", err)
	}

	checks := []struct {
		a   int64
		b   int64
		sum int64
	}{
		{3, 12, 15},
		{1, 4, 5},
		{16, 4, 3},
		{123, 623, 15},
	}

	for _, check := range checks {
		actual := gf.Add(big.NewInt(check.a), big.NewInt(check.b))
		if actual.Cmp(big.NewInt(check.sum)) != 0 {
			t.Errorf("%d + %d mod %d = %d; got %d", check.a, check.b, gf.P, check.sum, actual)
		}
	}
}

func TestSub(t *testing.T) {
	gf, err := NewGF(big.NewInt(17))
	if err != nil {
		t.Fatalf("Error generating finite field: %v", err)
	}

	checks := []struct {
		a    int64
		b    int64
		diff int64
	}{
		{3, 12, 8},
		{1, 4, 14},
		{16, 4, 12},
		{123, 623, 10},
	}

	for _, check := range checks {
		actual := gf.Sub(big.NewInt(check.a), big.NewInt(check.b))
		if actual.Cmp(big.NewInt(check.diff)) != 0 {
			t.Errorf("Expected %d - %d mod %d = %d; got %d", check.a, check.b, gf.P, check.diff, actual)
		}
	}
}

func TestMul(t *testing.T) {
	gf, err := NewGF(big.NewInt(17))
	if err != nil {
		t.Errorf("Error while creating new GF of prime order: %v", err)
	}

	checks := []struct {
		a    int64
		b    int64
		prod int64
	}{
		{1, 12, 12},
		{5, 4, 3},
		{12, 9, 6},
		{210, 152, 11},
	}

	for _, check := range checks {
		actual := gf.Mul(big.NewInt(check.a), big.NewInt(check.b))
		if actual.Cmp(big.NewInt(check.prod)) != 0 {
			t.Errorf("%d * %d mod %d = %d; got %d", check.a, check.b, gf.P, check.prod, actual)
		}
	}
}

func TestDiv(t *testing.T) {
	gf, err := NewGF(big.NewInt(17))
	if err != nil {
		t.Errorf("Error while creating new GF of prime order: %v", err)
	}

	checks := []struct {
		a    int64
		b    int64
		quot int64
	}{
		{1, 12, 12},
		{5, 4, 3},
		{12, 9, 6},
		{210, 152, 11},
	}

	for _, check := range checks {
		actual := gf.Div(big.NewInt(check.a), big.NewInt(check.b))
		if actual.Cmp(big.NewInt(check.quot)) != 0 {
			t.Errorf("%d / %d mod %d = %d; got %d", check.a, check.b, gf.P, check.quot, actual)
		}
	}
}

func TestMultInverse(t *testing.T) {
	gf, err := NewGF(big.NewInt(17))
	if err != nil {
		t.Errorf("Error while creating new GF of prime order: %v", err)
	}

	checks := []struct {
		a   int64
		inv int64
	}{
		{12, 10},
		{4, 13},
		{9, 2},
		{152, 16},
	}

	for _, check := range checks {
		actual := gf.MultInverse(big.NewInt(check.a))
		if actual.Cmp(big.NewInt(check.inv)) != 0 {
			t.Errorf("%d^-1 = %d^-1 mod %d; got %d", check.a, check.inv, gf.P, actual)
		}
	}
}

func TestExp(t *testing.T) {
	gf, err := NewGF(big.NewInt(17))
	if err != nil {
		t.Errorf("Error while creating new GF of prime order: %v", err)
	}

	checks := []struct {
		a   int64
		b   int64
		pow int64
	}{
		{1, 12, 1},
		{5, 4, 13},
		{12, 9, 5},
		{210, 152, 16},
	}

	for _, check := range checks {
		actual := gf.Exp(big.NewInt(check.a), big.NewInt(check.b))
		if actual.Cmp(big.NewInt(check.pow)) != 0 {
			t.Errorf("%d^%d mod %d = %d; got %d", check.a, check.b, gf.P, check.pow, actual)
		}
	}
}

func TestRand(t *testing.T) {
	gf, err := NewGF(big.NewInt(17))
	if err != nil {
		t.Errorf("Error while creating new GF of prime order: %v", err)
	}
	zero := big.NewInt(0)

	// It's probabilistic, but hey
	for i := 0; i < 10; i++ {
		rnd, err := gf.Rand()

		if err != nil {
			t.Errorf("Rand() returned error: %v", err)
		}

		if rnd.Cmp(zero) == -1 || rnd.Cmp(gf.P) != -1 {
			t.Errorf("Rand() = %d; Not valid for GF(%d)", rnd, gf.P)
		}
	}
}

func TestRandPolyonmial(t *testing.T) {
	gf, err := NewGF(big.NewInt(17))
	if err != nil {
		t.Errorf("Error while creating new GF of prime order: %v", err)
	}
	zero := big.NewInt(0)

	poly, err := gf.RandomPolynomial(3)

	if err != nil {
		t.Errorf("Polynomial generation failed: %v", err)
	}

	if len(poly.Coefficients) != 4 {
		t.Errorf("Expected polynomial of degree 3 to have 4  coefficients; got %d", len(poly.Coefficients))
	}

	for _, coef := range poly.Coefficients {
		if coef.Cmp(zero) == -1 || coef.Cmp(gf.P) != -1 {
			t.Errorf("Polynomial had coefficient which is not a group element: %d", coef)
		}
	}

	_, err = gf.RandomPolynomial(-1)
	if err == nil {
		t.Errorf("Expected error when generating polynomial of invalid degree; got none")
	}
}

func TestIsGroupElement(t *testing.T) {
	gf, err := NewGF(big.NewInt(17))
	if err != nil {
		t.Errorf("Error while creating new GF of prime order: %v", err)
	}

	valid := []int64{0, 2, 11, 16}
	for _, x := range valid {
		if !gf.IsGroupElement(big.NewInt(x)) {
			t.Errorf("Expected %d to be group element; was not", x)
		}
	}

	invalid := []int64{-3, 17, 256}
	for _, x := range invalid {
		if gf.IsGroupElement(big.NewInt(x)) {
			t.Errorf("Expected %d to not be group element; but was", x)
		}
	}
}
