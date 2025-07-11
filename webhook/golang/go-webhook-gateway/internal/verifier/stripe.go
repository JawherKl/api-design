package verifier

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strings"
	"time"
	"strconv"
)

func VerifyStripeSignature(signatureHeader string, payload []byte) bool {
	secret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	if secret == "" {
		return false
	}

	var timestamp, signature string

	parts := strings.Split(signatureHeader, ",")
	for _, part := range parts {
		if strings.HasPrefix(part, "t=") {
			timestamp = strings.TrimPrefix(part, "t=")
		}
		if strings.HasPrefix(part, "v1=") {
			signature = strings.TrimPrefix(part, "v1=")
		}
	}

	if timestamp == "" || signature == "" {
		return false
	}

	signedPayload := timestamp + "." + string(payload)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(signedPayload))
	expectedMAC := mac.Sum(nil)
	expectedSignature := hex.EncodeToString(expectedMAC)

	ts, _ := strconv.ParseInt(timestamp, 10, 64)
	if time.Now().Unix()-ts > 300 {
		return false
	}

	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}
