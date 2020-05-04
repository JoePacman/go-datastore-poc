package datastore_film

import (
	"cloud.google.com/go/datastore"
	"context"
	"go-datastore-poc/dto"
	"google.golang.org/api/iterator"
)

type Repo interface {
	FindByTitle(ctx context.Context, title string) ([]dto.Film, error)
	FindByPrimaryKey(ctx context.Context, pKey string) (*dto.Film, error)
	Update(ctx context.Context, Film *dto.Film) error
	Create(ctx context.Context, Film *dto.Film) error
}

type DatastoreRepo struct {
	client *datastore.Client
}

func NewDatastoreRepo(client *datastore.Client) *DatastoreRepo {
	return &DatastoreRepo{client: client}
}

func (repo *DatastoreRepo) FindByPrimaryKey(ctx context.Context, pKey string) (*dto.Film, error) {
	Film := new(dto.Film)

	if err := repo.client.Get(ctx, datastore.NameKey(dto.Kind, pKey, nil), Film); err != nil {
		return nil, err
	}
	return Film, nil
}

func (repo *DatastoreRepo) FindByTitle(ctx context.Context, title string) ([]dto.Film, error) {
	var films = make([]dto.Film, 0)

	query := datastore.NewQuery(dto.Kind).Filter("title=", title)

	it := repo.client.Run(ctx, query)

	for {
		var film dto.Film
		// _ underscore indicates an unused variable
		_, err := it.Next(&film)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		films = append(films, film)
	}
	return films, nil
}

func (repo *DatastoreRepo) Update(ctx context.Context, Film *dto.Film) error {
	panic("implement me")
}

func (repo *DatastoreRepo) Create(ctx context.Context, Film *dto.Film) error {
	//TODO key should not be title. Change out with MD5Hex
	k := datastore.NameKey(dto.Kind, Film.Title, nil)
	_, err := repo.client.Put(ctx, k, Film)
	if err != nil {
		return err
	} else {
		return nil
	}
}
