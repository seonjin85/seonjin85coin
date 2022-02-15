package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"

	"github.com/seonjin85/seonjin85coin/utils"
)

// 1) we hash the msg.
// "i love you" -> hash(x) -> "hashed_message"

// 2) generate key pair
// KeyPair ( privateK, publicK) (save privateK to a file)

// 3) sigh the hash
// ("hashed_message" + privateK) -> "signature"

// 4) verify
// ("hashed_message" + "signature" + publick) -> true / false

const (
	fileName string = "seonjin85.wallet"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	address    string
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func createPrivKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	utils.HandleErr(err)
	return privKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)
}

func restoreKey() *ecdsa.PrivateKey {
	keyAsBytes, err := os.ReadFile(fileName)
	utils.HandleErr(err)
	key, err := x509.ParseECPrivateKey(keyAsBytes)
	utils.HandleErr(err)
	return key
}

func aFromK(key *ecdsa.PrivateKey) string {

}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			w.privateKey = restoreKey()
		} else {
			key := createPrivKey()
			persistKey(key)
			w.privateKey = key
		}
		w.address = aFromK(w.privateKey)
	}
	return w
}
