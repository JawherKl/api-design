package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/JawherKl/go-webhook-gateway/internal/verifier"
	"log"
	"net/http"
	"io"
	"github.com/JawherKl/go-webhook-gateway/models"
	"time"
	"github.com/JawherKl/go-webhook-gateway/internal/storage"
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



	// Log the webhook event
	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		headers[k] = v[0]
	}

	event := models.WebhookEvent{
		Source:     source,
		ReceivedAt: time.Now(),
		Headers:    headers,
		Payload:    payload,
	}

	collection := storage.GetCollection("events")
	if _, err := collection.InsertOne(c, event); err != nil {
		log.Println("❌ Failed to store event:", err)
	}

	log.Printf("📦 Webhook stored from %s\n", source)
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