package handler

import (
	"context"
	"kafka-sarama-try/internal/domain"
	middleware "kafka-sarama-try/internal/middleware/jwt"
	"kafka-sarama-try/internal/usecase"
	"kafka-sarama-try/internal/util"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func OrderHandler(c *gin.Context) {
	authHeader := c.GetHeader("authorization")
	kontek := context.WithValue(c.Request.Context(), "waktu", time.Now())

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return
	}

	claims, err := middleware.VerifyJWT(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var logError error
	var logMessage string
	var logStatus int

	defer func() {
		if logError != nil {
			util.LogFailed(logMessage, c.Request.Method, kontek.Value("waktu").(time.Time), logStatus, logError)
		} else {
			util.LogSuccess(logMessage, c.Request.Method, kontek.Value("waktu").(time.Time), logStatus)
		}
	}()

	var orderReq domain.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// get the username of the person from their token
	orderReq.UserId = claims.Username

	log.Printf("===============================================================================================")
	log.Printf("Transaction ID %s for order type '%s' is STARTED\n", orderReq.TransactionID, orderReq.OrderType)
	log.Printf("===============================================================================================")

	_, _, err = usecase.OrderUsecase(&orderReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order": orderReq})
}
