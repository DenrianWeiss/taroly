package hx

import (
	"encoding/hex"
	"strings"
)

func HexStringToBytes(s string) []byte {
	s = strings.TrimPrefix(s, "0x")
	v, _ := hex.DecodeString(s)
	return v
}

func IsValidHex(s string) bool {
	s = strings.TrimPrefix(s, "0x")
	_, err := hex.DecodeString(s)
	return err == nil
}
