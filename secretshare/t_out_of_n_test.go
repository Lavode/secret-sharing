package secretshare

import (
	"github.com/lavode/secret-sharing/gf"
	"math/big"
	"testing"
)

func TestTOutOfN(t *testing.T) {
	field := gf.GF{P: big.NewInt(53)}
	secret := big.NewInt(42)
	t_ := 3
	n := 5

	shares, pol, err := TOutOfN(secret, t_, n, field)
	if err != nil {
		t.Fatalf("Error creating t-out-of-n share: %v", err)
	}

	if pol.Degree() != 2 {
		t.Errorf("Expected polynomial of degree 2; got %d", pol.Degree())
	}

	if len(shares) != 5 {
		t.Errorf("Expected 5 shares; got %d", len(shares))
	}

	_, _, err = TOutOfN(secret, 1, n, field)
	if err == nil {
		t.Errorf("Expected error if t <= 1; got none")
	}

	_, _, err = TOutOfN(secret, n, n, field)
	if err == nil {
		t.Errorf("Expected error if t >= n; got none")
	}

	_, _, err = TOutOfN(secret, 54, n, field)
	if err == nil {
		t.Errorf("Expected error if t not a group element; got none")
	}

	_, _, err = TOutOfN(secret, t_, 54, field)
	if err == nil {
		t.Errorf("Expected error if n not a group element; got none")
	}

	_, _, err = TOutOfN(big.NewInt(55), t_, n, field)
	if err == nil {
		t.Errorf("Expected error if secret not a group element; got none")
	}
}

func TestTOutOfNRecover(t *testing.T) {
	field := gf.GF{P: big.NewInt(53)}
	secret := big.NewInt(42)

	shares := []Share{
		{1, big.NewInt(37)},
		{2, big.NewInt(48)},
		{5, big.NewInt(18)},
	}
	actual, err := TOutOfNRecover(shares, field)
	if err != nil {
		t.Fatalf("Error while recovering secret: %v", err)
	}
	if actual.Cmp(secret) != 0 {
		t.Errorf("Expected to recover %d; got %d", secret, actual)
	}

	shares = []Share{
		{1, big.NewInt(37)},
		{3, big.NewInt(22)},
		{4, big.NewInt(12)},
	}
	actual, err = TOutOfNRecover(shares, field)
	if err != nil {
		t.Fatalf("Error while recovering secret: %v", err)
	}
	if actual.Cmp(secret) != 0 {
		t.Errorf("Expected to recover %d; got %d", secret, actual)
	}

	field = gf.GF{P: big.NewInt(127)}
	secret = big.NewInt(86)
	shares = []Share{
		{1, big.NewInt(30)},
		{3, big.NewInt(101)},
		{5, big.NewInt(109)},
		{6, big.NewInt(35)},
		{9, big.NewInt(86)},
	}
	actual, err = TOutOfNRecover(shares, field)
	if err != nil {
		t.Fatalf("Error while recovering secret: %v", err)
	}
	if actual.Cmp(secret) != 0 {
		t.Errorf("Expected to recover %d; got %d", secret, actual)
	}
}
