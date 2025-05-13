package repository

import (
	"errors"
	"github.com/looksaw/go_greenlight/internal/types"
)

type Repository interface {
	InsertIntoMovie(m types.MovieDto) (types.MovieDto, error)
	SelectMovieById(id int) (types.MovieDto, error)
	SelectMovieAll() ([]types.MovieDto, error)
	UpdateMovieById(id int, m types.MovieDto) (types.MovieDto, error)
	DeleteMovieById(id int)
}

// 下面是InMem的实现
type InMemRepository struct {
	Movies []types.MovieDto
}

func (r *InMemRepository) InsertIntoMovie(m types.MovieDto) (types.MovieDto, error) {
	r.Movies = append(r.Movies, m)
	return m, nil
}

func (r *InMemRepository) SelectMovieById(id int) (types.MovieDto, error) {
	n := len(r.Movies)
	if id >= n || id < 0 {
		return types.MovieDto{}, errors.New("id超过范围")
	}
	return r.Movies[id], nil
}

func (r *InMemRepository) SelectMovieAll() ([]types.MovieDto, error) {
	return r.Movies, nil
}
func (r *InMemRepository) UpdateMovieById(id int, m types.MovieDto) (types.MovieDto, error) {
	r.Movies[id] = m
	return m, nil
}
func (r *InMemRepository) DeleteMovieById(id int) {
	r.Movies[id] = types.MovieDto{}
}
