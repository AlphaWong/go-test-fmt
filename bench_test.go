package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func BenchmarkFmtConvertHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmtConvertHex(getDummySecretMessage())
	}
}

func BenchmarkEncodeHexConvertHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeHexConvertHex(getDummySecretMessage())
	}
}

func fmtConvertHex(secret []byte) string {
	// https://golang.org/pkg/crypto/sha256/#New
	return fmt.Sprintf("%x", secret)
}

func encodeHexConvertHex(secret []byte) string {
	return hex.EncodeToString(secret)
}

func getDummySecretMessage() []byte {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	return h.Sum(nil)
}
