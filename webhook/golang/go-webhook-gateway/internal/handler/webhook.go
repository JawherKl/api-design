package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/JawherKl/go-webhook-gateway/internal/verifier"
	"log"
	"net/http"
	"io"
)

func HandleWebhook(c *gin.Context) {
	source := c.Param("source")

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read body"})
		return
	}
	c.Request.Body = io.NopCloser(NewResetReader(bodyBytes))

	switch source {
	case "github":
		signature := c.GetHeader("X-Hub-Signature-256")
		if !verifier.VerifyGitHubSignature(signature, bodyBytes) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid GitHub signature"})
			return
		}
		log.Println("✅ GitHub signature verified")

	case "stripe":
		signature := c.GetHeader("Stripe-Signature")
		if !verifier.VerifyStripeSignature(signature, bodyBytes) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Stripe signature"})
			return
		}
		log.Println("✅ Stripe signature verified")

	case "twilio":
		signature := c.GetHeader("X-Twilio-Signature")
		if !verifier.VerifyTwilioSignature(signature, c.Request.URL.String(), c.Request.PostForm) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Twilio signature"})
			return
		}
		log.Println("✅ Twilio signature verified")	

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported webhook source"})
		return
	}

	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println("❌ Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	log.Printf("📦 Webhook received from %s: %+v\n", source, payload)
	c.JSON(http.StatusOK, gin.H{"status": "received", "source": source})
}

type ResetReader struct {
	data []byte
}

func NewResetReader(data []byte) *ResetReader {
	return &ResetReader{data: data}
}

func (r *ResetReader) Read(p []byte) (int, error) {
	copy(p, r.data)
	if len(r.data) == 0 {
		return 0, io.EOF
	}
	n := copy(p, r.data)
	r.data = r.data[n:]
	return n, nil
}
