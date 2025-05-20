package routes

import (
	"github.com/gin-gonic/gin"
	"week_9_crud/internal/auth"
	"week_9_crud/internal/db"
	"week_9_crud/internal/delivery"
	"week_9_crud/internal/middleware"
	"week_9_crud/internal/repository"
	"week_9_crud/internal/service"
)

func SetupRoutes(r *gin.Engine) {
	gameRepo := repository.NewGameRepository(db.DB)
	gameService := service.NewGameService(gameRepo)
	gameHandler := delivery.NewGameHandler(gameService)

	games := r.Group("api/games")
	{
		games.GET("/", gameHandler.GetAllGames)
		games.GET("/:id", gameHandler.GetGameById)

		// Требуют авторизации
		games.POST("/", middleware.AuthRequired(), gameHandler.CreateGame)
		games.PUT("/:id", middleware.AuthRequired(), gameHandler.UpdateGame)
		games.DELETE("/:id", middleware.AuthRequired(), gameHandler.DeleteGame)
	}

	reviewRepo := repository.NewReviewRepository(db.DB)
	reviewService := service.NewReviewService(reviewRepo)
	reviewHandler := delivery.NewReviewHandler(reviewService)

	reviews := r.Group("/api/games/:id/reviews")
	{
		reviews.GET("/", reviewHandler.GetReviewsByGameID)
		reviews.POST("/", middleware.AuthRequired(), reviewHandler.CreateReview)
	}

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST("/login", auth.Login)
	}

	userRepo := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepo)
	userHandler := delivery.NewUserHandler(userService)

	userRoutes := r.Group("api/users")
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.GET("/me", middleware.AuthRequired(), auth.GetCurrentUser)
	}

}
