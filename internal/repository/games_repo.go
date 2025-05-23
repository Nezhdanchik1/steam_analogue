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

func (r GameRepositoryImpl) GetAllGames() ([]models.Game, error) {
	var games []models.Game
	err := r.db.Find(&games).Error
	return games, err
}

func (r GameRepositoryImpl) GetGameById(id int) (*models.Game, error) {
	var game models.Game
	err := r.db.First(&game, id).Error
	return &game, err
}

func (r GameRepositoryImpl) CreateGame(game *models.Game) error {
	return r.db.Create(game).Error
}

func (r GameRepositoryImpl) UpdateGame(id int, game *models.GameEdit) error {
	return r.db.Model(&models.Game{}).Where("id = ?", id).Omit("id").Updates(game).Error
}

func (r GameRepositoryImpl) DeleteGame(id int) error {
	return r.db.Delete(&models.Game{}, id).Error
}
