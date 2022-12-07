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

func FilterUnPrintable(s string) string {
	r := ""
	for _, c := range s {
		if c >= 32 && c <= 126 {
			r += string(c)
		}
	}
	return r
}

func IsValidHex(s string) bool {
	s = strings.TrimPrefix(s, "0x")
	_, err := hex.DecodeString(s)
	return err == nil
}
