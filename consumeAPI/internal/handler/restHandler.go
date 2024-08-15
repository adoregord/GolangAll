package handler

import (
	"consume-api-try/internal/domain"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestHandler struct{}

type RestHandlerInterface interface {
	ConsumeRest(c *gin.Context)
}

func NewRestHandler() RestHandlerInterface {
	return RestHandler{}
}

func (h RestHandler) ConsumeRest(c *gin.Context) {
	resp, err := http.Get("http://localhost:8080/user/get")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer resp.Body.Close()

	var user []domain.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
