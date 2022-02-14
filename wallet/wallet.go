package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

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
	signature     string = "f32baeb3e88f64995cb7fcf47f5609b823e4cf31b9c8a9f85840ac54b7dbba21ae518767c244f0c9d415d91e7a1abbb8ba64910b9809f52be82dca5eac145e42"
	privateKey    string = "307702010104208ef4e9f443c6ef655a209b7630c2f5cf65be52a79eeb06d062aae856a464ac5aa00a06082a8648ce3d030107a14403420004d0cbe458f8f0a56f3851b86e8119e3d73b5af0c6c27b5f8e6e3d56147a4ee1d8525e9a60a9a5839c697fe55f8932ede8927874e15e944e827a24c66353a2970b"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	privateKey, err = x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
}
