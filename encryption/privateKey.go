package encryption

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"math/rand"
)

type PrivateKey struct {
	*PublicKey
	D *big.Int
}

func GenerateKeyPair() (*PrivateKey, error) {
	curve := getCurve()

	p, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("could not generate key pair: %v", err)
	}

	return &PrivateKey{
		PublicKey: &PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(p),
	}, nil
}

func getCurve() elliptic.Curve {
	return elliptic.P256()
}
