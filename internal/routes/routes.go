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

	// Группа маршрутов для игр
	games := r.Group("api/games")
	{
		games.GET("/", gameHandler.GetAllGames)
		games.GET("/:id", gameHandler.GetGameById)
		games.POST("/", gameHandler.CreateGame)
		games.PUT("/:id", gameHandler.UpdateGame)
		games.DELETE("/:id", gameHandler.DeleteGame)
	}

	// Группа маршрутов для авторизации
	authRoutes := r.Group("api/auth")
	{
		// authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
		authRoutes.POST("/login", auth.Login)
	}

	userRepo := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepo)
	userHandler := delivery.NewUserHandler(userService)
	// Группа маршрутов для пользователей
	userRoutes := r.Group("api/users")
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
	}

	protected := r.Group("api/users")
	protected.Use(middleware.AuthRequired())
	protected.Use(middleware.RoleMiddleware("admin"))
	{
		protected.GET("/me", auth.GetCurrentUser)
	}

}
