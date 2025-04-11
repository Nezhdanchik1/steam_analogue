package main

import (
	"github.com/gin-gonic/gin"
	"week_9_crud/internal/db"
	"week_9_crud/internal/routes"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
