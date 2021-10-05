package password

import (
	"fmt"
	"math/rand"
	"time"
)

func randomString(length int8) string {
    rand.Seed(time.Now().UnixNano())
    bytes := make([]byte, length)
    rand.Read(bytes)
    return fmt.Sprintf("%x", bytes)[:length]
}

func Generate(length int8) string {
	return randomString(length)
}
