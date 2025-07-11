package verifier

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func VerifyGitHubSignature(signatureHeader string, body []byte) bool {
	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")
	if secret == "" {
		fmt.Println("❌ Missing GITHUB_WEBHOOK_SECRET")
		return false
	}

	expectedPrefix := "sha256="
	if len(signatureHeader) <= len(expectedPrefix) || signatureHeader[:len(expectedPrefix)] != expectedPrefix {
		fmt.Println("❌ Signature header format invalid")
		return false
	}

	signature := signatureHeader[len(expectedPrefix):]

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedMAC := mac.Sum(nil)
	expectedSignature := hex.EncodeToString(expectedMAC)

	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}
