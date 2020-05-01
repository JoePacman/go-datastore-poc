package datastore_film

import (
	"context"
	"errors"
	"go-datastore-poc/dto"
)

type FilmService struct {
	repo Repo
}

func NewService(repo Repo) *FilmService {
	return &FilmService{repo: repo}
}

// These methods (belongs to the service), not a function;
// "(service *FilmService)" says that this method belongs to a "FilmService" struct
func (service *FilmService) FindByTitle(ctx context.Context, name string) ([]dto.Film, error) {
	film, err := service.repo.FindByTitle(ctx, name)

	if err != nil {
		return nil, errors.New("error fetching Film; " + err.Error())
	}
	return film, nil
}

func(service *FilmService) Create(ctx context.Context, film *dto.Film) error {
	return service.repo.Create(ctx, film)
}
