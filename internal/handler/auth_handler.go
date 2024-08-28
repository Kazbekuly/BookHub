package handler

import (
	"BookHub/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUP(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		return
	}
	id, err := h.services.CreateUser(user)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type SignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input SignIn
	if err := c.BindJSON(&input); err != nil {
		return
	}
	token, err := h.services.GenerateToken(input.Username, input.Password)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})

}
