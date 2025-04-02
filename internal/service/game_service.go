package service

import "week_9_crud/internal/models"

type GameRepository interface {
	GetAllGames() ([]models.Game, error)
	GetGameById(id int) (*models.Game, error)
	CreateGame(game *models.Game) error
	UpdateGame(id int, game *models.GameEdit) error
	DeleteGame(id int) error
}

type GameService struct {
	repo GameRepository
}

func NewGameService(gameRepo GameRepository) *GameService {
	return &GameService{repo: gameRepo}
}

func (g *GameService) GetAllGames() ([]models.Game, error) {
	return g.repo.GetAllGames()
}

func (g *GameService) GetGameById(id int) (*models.Game, error) {
	return g.repo.GetGameById(id)
}

func (g *GameService) CreateGame(title, description, developer string, price float64) (*models.Game, error) {
	game := &models.Game{
		Title:       title,
		Description: description,
		Developer:   developer,
		Price:       price,
	}
	err := g.repo.CreateGame(game)
	return game, err
}

func (g *GameService) UpdateGame(id int, gameEdit *models.GameEdit) (*models.Game, error) {
	err := g.repo.UpdateGame(id, gameEdit)
	if err != nil {
		return nil, err
	}
	return g.GetGameById(id)
}

func (g *GameService) DeleteGame(id int) error {
	return g.repo.DeleteGame(id)
}
