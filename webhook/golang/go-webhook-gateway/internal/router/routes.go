package router

import (
	"github.com/gin-gonic/gin"
	"github.com/JawherKl/go-webhook-gateway/internal/handler"
)

func SetupRoutes(r *gin.Engine) {
	webhook := r.Group("/webhook")
	{
		webhook.POST("/:source", handler.HandleWebhook)
	}
}
