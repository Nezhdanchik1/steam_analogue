package repository

import (
	"gorm.io/gorm"
	"week_9_crud/internal/models"
)

type GameRepositoryImpl struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GameRepositoryImpl {
	return &GameRepositoryImpl{db: db}
}

func (g GameRepositoryImpl) GetAllGames() ([]models.Game, error) {
	var games []models.Game
	err := g.db.Find(&games).Error
	return games, err
}
