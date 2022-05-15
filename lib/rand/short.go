package rand

import (
	"crypto/rand"
	"encoding/hex"
)

func ShortString(length int) string {
	b := make([]byte, length)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
