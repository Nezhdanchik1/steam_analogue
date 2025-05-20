package service

import (
	"week_9_crud/internal/models"
)

type ReviewRepository interface {
	GetReviewsByGameID(gameID uint) ([]models.Review, error)
	CreateReview(review *models.Review) error
}

type ReviewService struct {
	repo ReviewRepository
}

func NewReviewService(repo ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) GetReviewsByGameID(gameID uint) ([]models.Review, error) {
	return s.repo.GetReviewsByGameID(gameID)
}

func (s *ReviewService) CreateReview(review *models.Review) error {
	return s.repo.CreateReview(review)
}
