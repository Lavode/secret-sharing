package gf

import (
	"math/big"
	"testing"
)

func TestBasePolynomial(t *testing.T) {
	gf := GF{P: big.NewInt(53)}
	xs := []*big.Int{
		big.NewInt(1),
		big.NewInt(3),
		big.NewInt(5),
	}

	actual := BasePolynomial(0, xs, gf)
	if actual.Cmp(big.NewInt(35)) != 0 {
		t.Errorf("Expected l_0(0) = 35; got %d", actual)
	}

	actual = BasePolynomial(1, xs, gf)
	if actual.Cmp(big.NewInt(12)) != 0 {
		t.Errorf("Expected l_1(0) = 12; got %d", actual)
	}

	actual = BasePolynomial(2, xs, gf)
	if actual.Cmp(big.NewInt(7)) != 0 {
		t.Errorf("Expected l_2(0) = 7; got %d", actual)
	}
}
