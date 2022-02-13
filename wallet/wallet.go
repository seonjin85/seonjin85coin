package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

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

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	fmt.Printf("privateKey:%d\n", privateKey.D)
	fmt.Printf("publickey:\nx:%d,\ny:%d\n", privateKey.X, privateKey.Y)

	message := "i love you"
	hashedMessage := utils.Hash(message)
	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleErr(err)
	fmt.Printf("R:%d\nS:%d", r, s)
}
