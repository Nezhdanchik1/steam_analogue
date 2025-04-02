package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"week_9_crud/internal/models"
	"week_9_crud/internal/routes"
)
import "gorm.io/driver/postgres"

func main() {
	db, err := gorm.Open(postgres.Open("postgres://nezhdanchik:qwerty@localhost:5433/games_DB?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&models.Game{})
	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8080")
}
