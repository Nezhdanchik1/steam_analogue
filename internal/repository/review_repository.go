package repository

import (
	"gorm.io/gorm"
	"week_9_crud/internal/models"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) GetReviewsByGameID(gameID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("game_id = ?", gameID).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepository) CreateReview(review *models.Review) error {
	return r.db.Create(review).Error
}
