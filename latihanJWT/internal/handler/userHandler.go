package handler

import (
	"encoding/json"
	"jwt-try/internal/domain"
	"jwt-try/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	RegisterUser(c *gin.Context)
	CheckCredential(c *gin.Context)
	ViewWeb(c *gin.Context)
	CheckVerifiedAcc(c *gin.Context)
}

type UserHandler struct {
	UserUsecase usecase.UserUsecaseInterface
}

func NewUserHandler(userUsecase usecase.UserUsecaseInterface) UserHandlerInterface {
	return UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h UserHandler) RegisterUser(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	// check if the method is post
	if c.Request.Method != "POST" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
		return
	}

	var user domain.User

	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	User, err := h.UserUsecase.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"UserID": User.ID, "Username": User.Username})
}

func (h UserHandler) CheckCredential(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	// check if the method is post
	if c.Request.Method != "POST" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
		return
	}

	var user domain.User

	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.UserUsecase.CheckCredential(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token})
}

func (h UserHandler) ViewWeb(c *gin.Context) {
	authHeader := c.GetHeader("AUTH")
	c.Writer.Header().Set("Content-Type", "application/json")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return
	}

	claims, err := h.UserUsecase.VerifyJWT(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Welcome !!!",
		Data:    claims,
	})
}

func (h UserHandler) CheckVerifiedAcc(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	// check if the method is post
	if c.Request.Method != "POST" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
		return
	}

	type req struct {
		Username string
	}

	var a req

	if err := json.NewDecoder(c.Request.Body).Decode(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, err := h.UserUsecase.CheckVerifiedAcc(a.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if ok {
		c.JSON(http.StatusAccepted, gin.H{"Status": "Account is Verified"})
		return
	}

	c.JSON(400, gin.H{"Status": "Account is Not Verified"})
}
