package routes

import (
	"github.com/gin-gonic/gin"
	"week_9_crud/internal/db"
	"week_9_crud/internal/delivery"
	"week_9_crud/internal/repository"
	"week_9_crud/internal/service"
)

func SetupRoutes(r *gin.Engine) {
	gameRepo := repository.NewGameRepository(db.DB)

	gameService := service.NewGameService(gameRepo)

	gameHandler := delivery.NewGameHandler(gameService)

	// Настроим маршруты
	games := r.Group("api/games")
	{
		games.GET("/", gameHandler.GetAllGames)
		games.GET("/:id", gameHandler.GetGameById)
		games.POST("/", gameHandler.CreateGame)
		games.PUT("/:id", gameHandler.UpdateGame)
		games.DELETE("/:id", gameHandler.DeleteGame)
	}
}
