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

func (g GameRepositoryImpl) GetGameById(id int) (*models.Game, error) {
	var game models.Game
	err := g.db.First(&game, id).Error
	return &game, err
}

func (g GameRepositoryImpl) CreateGame(game *models.Game) error {
	return g.db.Create(game).Error
}

func (g GameRepositoryImpl) UpdateGame(id int, game *models.GameEdit) error {
	return g.db.Model(&models.Game{}).Where("id = ?", id).Omit("id, CreatedAt").Updates(game).Error
}

func (g GameRepositoryImpl) DeleteGame(id int) error {
	return g.db.Delete(&models.Game{}, id).Error
}
