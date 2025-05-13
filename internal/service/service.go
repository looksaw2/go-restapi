package service

import (
	"net/http"

	"github.com/looksaw/go_greenlight/internal/repository"
	"github.com/looksaw/go_greenlight/internal/types"
)

type Service interface {
	CreateMovie(req types.CreateMovieRequest) types.CreateMovieResponse
	ShowMovieById(id int) (types.CreateMovieResponse, error)
	ShowMovieAll() []types.CreateMovieResponse
	UpdateMovieById(id int, req types.CreateMovieRequest) (types.CreateMovieResponse, error)
	DeleteMovieById(id int) types.DeleteResponse
}

type InMemService struct {
	repository repository.Repository
}

func NewInMemService(repository repository.Repository) *InMemService {
	return &InMemService{
		repository: repository,
	}
}

func (s *InMemService) CreateMovie(req types.CreateMovieRequest) types.CreateMovieResponse {
	res, _ := s.repository.InsertIntoMovie(*req.ToDto())
	return *res.ToResponse(http.StatusCreated, "Movie Created")
}

func (s *InMemService) ShowMovieById(id int) (types.CreateMovieResponse, error) {
	res, err := s.repository.SelectMovieById(id)
	if err != nil {
		return types.CreateMovieResponse{}, err
	}
	return *res.ToResponse(http.StatusOK, "Movie already found"), nil
}

func (s *InMemService) ShowMovieAll() []types.CreateMovieResponse {
	mk, err := s.repository.SelectMovieAll()
	if err != nil {
		return []types.CreateMovieResponse{}
	}
	var ans []types.CreateMovieResponse
	for _, m := range mk {
		ans = append(ans, *m.ToResponse(http.StatusOK, "Movie already found"))
	}
	return ans
}

func (s *InMemService) UpdateMovieById(id int, req types.CreateMovieRequest) (types.CreateMovieResponse, error) {
	res, err := s.repository.UpdateMovieById(id, *req.ToDto())
	if err != nil {
		return types.CreateMovieResponse{}, err
	}
	return *res.ToResponse(http.StatusOK, "Update Movie Success"), nil
}

func (s *InMemService) DeleteMovieById(id int) types.DeleteResponse {
	s.repository.DeleteMovieById(id)
	return types.DeleteResponse{
		DeleteID: id,
		Err:      nil,
	}
}
