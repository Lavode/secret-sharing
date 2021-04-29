package galois_field

import (
	"math/big"
	"testing"
)

func TestAdd(t *testing.T) {
	gf := GF{P: big.NewInt(17)}

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

func TestMul(t *testing.T) {
	gf := GF{P: big.NewInt(17)}

	checks := []struct {
		a   int64
		b   int64
		sum int64
	}{
		{1, 12, 12},
		{5, 4, 3},
		{12, 9, 6},
		{210, 152, 11},
	}

	for _, check := range checks {
		actual := gf.Mul(big.NewInt(check.a), big.NewInt(check.b))
		if actual.Cmp(big.NewInt(check.sum)) != 0 {
			t.Errorf("%d * %d mod %d = %d; got %d", check.a, check.b, gf.P, check.sum, actual)
		}
	}
}

func TestRand(t *testing.T) {
	gf := GF{P: big.NewInt(17)}
	min := big.NewInt(0)

	// It's probabilistic, but hey
	for i := 0; i < 10; i += 1 {
		rnd, err := gf.Rand()

		if err != nil {
			t.Errorf("Rand() returned error: %v", err)
		}

		if rnd.Cmp(min) == -1 || rnd.Cmp(gf.P) != -1 {
			t.Errorf("Rand() = %d; Not valid for GF(%d)", rnd, gf.P)
		}
	}
}
