package uuid

import (
	"encoding/hex"
	"github.com/google/uuid"
)

func NewAsHex() string {
	return hex.EncodeToString([]byte(uuid.New().String()))
}
