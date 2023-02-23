package repositories

import (
	"finaltaskbe/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindFilms() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	AddFilm(film models.Film) (models.Film, error)
	GetOneFilm(ID, User int) (models.Film, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFilms() ([]models.Film, error) {
	var films []models.Film
	err := r.db.Find(&films).Error
	return films, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
	var film models.Film
	err := r.db.First(&film, ID).Error

	return film, err
}

func (r *repository) GetOneFilm(ID, User int) (models.Film, error) {
	var film models.Film
	err := r.db.First(&film, ID).Error

	var trans models.Transaction
	error := r.db.Where("film_id = ? AND user_id = ?", ID, User).First(&trans).Error
	if error == nil && trans.Status == "success" {
		film.Price = 0
	}

	return film, err
}

func (r *repository) AddFilm(film models.Film) (models.Film, error) {
	err := r.db.Create(&film).Error

	return film, err
}

func (r *repository) DeleteFilm(film models.Film) (models.Film, error) {
	err := r.db.Delete(&film).Error

	return film, err
}
