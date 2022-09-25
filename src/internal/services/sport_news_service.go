package services

import (
	"context"
	"incrowd/src/internal/model"
	"incrowd/src/internal/ports"
)

type SportNewsService struct {
	relationalSportNewsDBRepository ports.NonRelationalSportNewsDBRepository
}

func NewSportNewsService(relationalSportNewsDBRepository ports.NonRelationalSportNewsDBRepository) *SportNewsService {
	return &SportNewsService{
		relationalSportNewsDBRepository: relationalSportNewsDBRepository,
	}
}

func (s *SportNewsService) StoreNews(ctx context.Context, news []model.News) error {
	return s.relationalSportNewsDBRepository.StoreNews(ctx, news)
}

func (s *SportNewsService) GetNewsWithID(ctx context.Context, id string) (*model.News, error) {
	return s.relationalSportNewsDBRepository.GetNewsWithID(ctx, id)
}

func (s *SportNewsService) GetNews(ctx context.Context) ([]model.News, error) {
	return s.relationalSportNewsDBRepository.GetNews(ctx)
}
