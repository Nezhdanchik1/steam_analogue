package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"week_9_crud/internal/service"
)

type GameHandler struct {
	service *service.GameService
}

func NewGameHandler(service *service.GameService) *GameHandler {
	return &GameHandler{service: service}
}

func (g *GameHandler) GetAllGames(c *gin.Context) {
	games, _ := g.service.GetAllGames()
	c.JSON(http.StatusOK, games)
}
