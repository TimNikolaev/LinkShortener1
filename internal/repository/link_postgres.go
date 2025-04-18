package repository

import (
	"shortener"

	"gorm.io/gorm"
)

type LinkPostgres struct {
	db *gorm.DB
}

func NewLinkPostgres(db *gorm.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func (r *LinkPostgres) Create(link *shortener.Link) (*shortener.Link, error) {
	result := r.db.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (r *LinkPostgres) GetByHash(hash string) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (r *LinkPostgres) GetByID(id uint) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (r *LinkPostgres) Update(link shortener.Link) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (r *LinkPostgres) Delete(id uint) error {
	return nil
}
