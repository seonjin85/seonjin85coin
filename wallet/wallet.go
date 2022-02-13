package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
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

const (
	signature     string = "f32baeb3e88f64995cb7fcf47f5609b823e4cf31b9c8a9f85840ac54b7dbba21ae518767c244f0c9d415d91e7a1abbb8ba64910b9809f52be82dca5eac145e42"
	privateKey    string = "307702010104208ef4e9f443c6ef655a209b7630c2f5cf65be52a79eeb06d062aae856a464ac5aa00a06082a8648ce3d030107a14403420004d0cbe458f8f0a56f3851b86e8119e3d73b5af0c6c27b5f8e6e3d56147a4ee1d8525e9a60a9a5839c697fe55f8932ede8927874e15e944e827a24c66353a2970b"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)
	utils.HandleErr(err)
	fmt.Printf("%x\n", keyAsBytes)

	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleErr(err)

	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("%x\n", signature)

}
