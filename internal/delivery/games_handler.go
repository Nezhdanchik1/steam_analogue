package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"week_9_crud/internal/models"
	"week_9_crud/internal/service"
)

type GameHandler struct {
	service *service.GameService
}

func NewGameHandler(service *service.GameService) *GameHandler {
	return &GameHandler{service: service}
}

func (h *GameHandler) GetAllGames(c *gin.Context) {
	games, _ := h.service.GetAllGames()
	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) GetGameById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	game, err := h.service.GetGameById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) CreateGame(c *gin.Context) {
	var gameCreate models.GameEdit

	if err := c.ShouldBindJSON(&gameCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newGame, err := h.service.CreateGame(gameCreate.Title, gameCreate.Description, gameCreate.Developer, gameCreate.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game"})
		return
	}

	c.JSON(http.StatusCreated, newGame)
}

func (h *GameHandler) UpdateGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	var gameEdit models.GameEdit
	if err := c.ShouldBindJSON(&gameEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedGame, err := h.service.UpdateGame(id, &gameEdit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, updatedGame)
}

func (h *GameHandler) DeleteGame(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
		return
	}

	if err := h.service.DeleteGame(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}
