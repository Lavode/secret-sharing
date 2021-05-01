package main

import (
	"fmt"
	"github.com/lavode/secret-sharing/gf"
	"github.com/lavode/secret-sharing/secretshare"
	"math/big"
	"os"
)

func main() {
	t := 3
	n := 5
	var q = &big.Int{}
	// 1024 bit prime
	q.SetString("162185254415496084206340421140209635914011817830718629561966397771320090737976933417510681112145040759352403578123408138464090606028577476738969675649310093809320506675661674944359979295818459722388210685500692240179292873664199942911683578488773858384553254621913954467555875395131791893173015885067312908733", 10)
	secret := big.NewInt(42)

	shares, err := share(secret, t, n, q)
	if err != nil {
		fmt.Printf("Secret sharing failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d-out-of-%d secret sharing of %d produced shares:\n", t, n, secret)
	for _, share := range shares {
		fmt.Printf("  Share %d = %d\n", share.ID, share.Value)
	}

	fmt.Println("Reconstructing secret using three shares")
	reconstructShares := []secretshare.Share{
		shares[0],
		shares[2],
		shares[4],
	}

	reconstructed, err := reconstruct(reconstructShares, q)
	if err != nil {
		fmt.Printf("Secret reconstruction failed: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Secret reconstructed: %d\n", reconstructed)
}

// share implements (t+1)-out-of-n secret sharing using polynomials of degree t
// over a finite field of prime order q.
//
// Returns a list of shares, or an error if share generation failed.
func share(secret *big.Int, t int, n int, q *big.Int) ([]secretshare.Share, error) {
	var shares []secretshare.Share

	field, err := gf.NewGF(q)
	if err != nil {
		return shares, err
	}

	shares, _, err = secretshare.TOutOfN(secret, t, n, field)
	return shares, err
}

// reconstruct implements secret recovery using t+1 shares. The list of
// supplied shares must contain exactly t+1 shares, otherwise operation will be
// incorrect.
//
// Returns the reconstructed secret, or an error if reconstruction failed.
func reconstruct(shares []secretshare.Share, q *big.Int) (*big.Int, error) {
	var secret *big.Int

	field, err := gf.NewGF(q)
	if err != nil {
		return secret, err
	}

	secret, err = secretshare.TOutOfNRecover(shares, field)
	return secret, err
}
