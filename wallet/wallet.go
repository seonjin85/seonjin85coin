package wallet

import (
	"crypto/ecdsa"
	"os"
)

// 1) we hash the msg.
// "i love you" -> hash(x) -> "hashed_message"

// 2) generate key pair
// KeyPair ( privateK, publicK) (save privateK to a file)

// 3) sigh the hash
// ("hashed_message" + privateK) -> "signature"

// 4) verify
// ("hashed_message" + "signature" + publick) -> true / false

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat("seonjin85.wallet")
	return !os.IsNotExist(err)
}

func Wallet() *wallet {
	if w == nil {

		if hasWalletFile() {

		}
	}
	return w
}
