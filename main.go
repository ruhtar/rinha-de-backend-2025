package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	r.POST("/payments", func(c *gin.Context) {
		var req PaymentRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":       "Pagamento recebido com sucesso",
			"correlationId": req.CorrelationID,
			"amount":        req.Amount,
		})
	})

	r.Run(":9999")
}

type PaymentRequest struct {
	CorrelationID string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
}
