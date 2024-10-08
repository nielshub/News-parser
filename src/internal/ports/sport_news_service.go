package ports

import (
	"context"
	"incrowd/src/internal/model"
)

type SportNewsService interface {
	StoreNews(context.Context, []model.News) error
	GetNewsWithID(context.Context, string) (*model.News, error)
	GetNews(context.Context) ([]model.News, error)
}
