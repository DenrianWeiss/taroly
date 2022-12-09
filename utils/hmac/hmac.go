package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/DenrianWeiss/taroly/utils/hx"
	"github.com/google/uuid"
)

var nonce string

func init() {
	nonce = uuid.New().String()
}

func GetNonce() string {
	return nonce
}

func SignWithNonce(s string) string {
	hmac := hmac.New(sha256.New, []byte(GetNonce()))
	hash := hmac.Sum([]byte(s))
	return hex.EncodeToString(hash)
}

func Validate(s, sig string) bool {
	sHash := SignWithNonce(s)
	sigHex := hx.HexStringToBytes(sig)
	return hmac.Equal([]byte(sHash), sigHex)
}
