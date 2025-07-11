package verifier

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"os"
	"sort"
	"strings"
)

func VerifyTwilioSignature(signatureHeader string, fullURL string, form url.Values) bool {
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	if authToken == "" {
		return false
	}

	var b strings.Builder
	b.WriteString(fullURL)

	keys := make([]string, 0, len(form))
	for k := range form {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		b.WriteString(key)
		b.WriteString(form.Get(key))
	}

	data := b.String()

	mac := hmac.New(sha1.New, []byte(authToken))
	mac.Write([]byte(data))
	expectedSig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(expectedSig), []byte(signatureHeader))
}
