package service

import "week_9_crud/internal/models"

type GameRepository interface {
	GetAllGames() ([]models.Game, error)
}

type GameService struct {
	repo GameRepository
}

func NewStudentService(gameRepo GameRepository) *GameService {
	return &GameService{repo: gameRepo}
}

func (g *GameService) GetAllGames() ([]models.Game, error) {
	return g.repo.GetAllGames()
}
