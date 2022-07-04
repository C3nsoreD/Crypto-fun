package encryption

import (
	"crypto/elliptic"
	"math/big"
)

type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}
