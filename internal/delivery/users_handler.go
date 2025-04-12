package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"week_9_crud/internal/service"
)

type UserHandler struct {
	s *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{s: s}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, _ := h.s.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
