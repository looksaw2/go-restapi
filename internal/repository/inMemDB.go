package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

type PGRepository interface {
	GetMovie(ctx context.Context, id int32) (*Movie, error)
	GetListMovie(ctx context.Context) ([]Movie, error)
	CreateMovie(ctx context.Context, name, email string) (*Movie, error)
	UpdateMovie(ctx context.Context, id int64, name, email string) (*Movie, error)
	DeleteMovie(ctx context.Context, id int64) error
}

// postrgresql的实现
type PostgresqlRepository struct {
	db    *pgx.Conn
	query *Queries
}

// postgresql初始化
func NewPostgresqlRepository(db *pgx.Conn) PostgresqlRepository {
	return PostgresqlRepository{
		db:    db,
		query: New(db),
	}
}

// Postgresql创建movie
func (r *PostgresqlRepository) GetMovie(ctx context.Context, id int64) (*Movie, error) {
	movie, err := r.query.GetMovie(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Movie{
		ID:    movie.ID,
		Name:  movie.Name,
		Email: movie.Email,
	}, nil
}

// Postgresql的得到所有的movie
func (r *PostgresqlRepository) getListMovie(ctx context.Context) ([]Movie, error) {
	var ans []Movie
	movies, err := r.query.GetListMovie(ctx)
	if err != nil {
		return nil, err
	}
	for _, movie := range movies {
		ans = append(ans, movie)
	}
	return ans, nil
}

// Postgresql创建新的movie
func (r *PostgresqlRepository) CreateMovie(ctx context.Context, name string, email string) (*Movie, error) {
	movie, err := r.query.CreateMovie(ctx, CreateMovieParams{
		Name:  pgtype.Text{String: name, Valid: true},
		Email: pgtype.Text{String: email, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &Movie{
		ID:    movie.ID,
		Name:  movie.Name,
		Email: movie.Email,
	}, nil
}

// Postgresql修改Movie
func (r *PostgresqlRepository) UpdateMovie(ctx context.Context, id int64, name string, email string) (*Movie, error) {
	movie, err := r.query.UpdateMovie(ctx, UpdateMovieParams{
		ID:    id,
		Name:  pgtype.Text{String: name, Valid: true},
		Email: pgtype.Text{String: email, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &Movie{
		ID:    movie.ID,
		Name:  movie.Name,
		Email: movie.Email,
	}, nil
}

// Postgresql删除Movie
func (r *PostgresqlRepository) DeleteMovie(ctx context.Context, id int64) error {
	err := r.query.DeleteMovie(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
