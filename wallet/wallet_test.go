package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"io/fs"
	"reflect"
	"testing"
)

const (
	testKey     string = "30770201010420bd6bd8d7c6607d92f012ec2b4941cb841a2f5efb43f7503cf6f2924d26daa47ca00a06082a8648ce3d030107a14403420004df5d824fae61aa80fec492d49d0dbdd920145fc9c662c4aea91a580593cffe30680de148e1c69ba285f6663d1bc305e035f4044a3dbf23e5513b9af549f4b404"
	testPayload string = "0000cd1f2c78cbc035fbe2292ba6765efc36eb9d7c2c1bdc73d6dc0d98386199"
	testSign    string = "30770201010420bd6bd8d7c6607d92f012ec2b4941cb841a2f5efb43f7503cf6f2924d26daa47ca00a06082a8648ce3d030107a14403420004df5d824fae61aa80fec492d49d0dbdd920145fc9c662c4aea91a580593cffe30680de148e1c69ba285f6663d1bc305e035f4044a3dbf23e5513b9af549f4b404"
)

type fakeLayer struct {
	fakeHasWalletFile func() bool
}

func (f fakeLayer) hasWalletFile() bool {
	return f.fakeHasWalletFile()
}

func (fakeLayer) writeFile(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func (fakeLayer) readFile(name string) ([]byte, error) {
	// return utils.ToBytes(makeTestWallet().privateKey), nil
	return x509.MarshalECPrivateKey(makeTestWallet().privateKey)
}

func TestWallet(t *testing.T) {
	t.Run("New Wallet is created", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool {
				t.Log("I have been called")
				return false
			},
		}
		w := Wallet()
		if reflect.TypeOf(w) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})

	t.Run("Wallet is restored", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool {
				t.Log("I have been called")
				return true
			},
		}
		w = nil
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New wallet should return a new wallet instance")
		}
	})
}

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}
	tests := []test{
		{testPayload, true},
		{"0055cd1f2c78cbc035fbe2292ba6765efc36eb9d7c2c1bdc73d6dc0d98386199", false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSign, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("verify() could not verify testSignature and Payload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("restroeBigInts should return error when payload is not hex.")
	}
}
