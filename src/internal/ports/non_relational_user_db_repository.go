package ports

import (
	"context"
	"incrowd/src/internal/model"
)

type NonRelationalSportsNewDBRepository interface {
	StoreNews(context.Context, []model.News) error
	GetNewsWithID(context.Context, string, string) (*model.News, error)
	GetNews(context.Context) ([]model.News, error)
}
