package types

// 传入的Movie请求
type CreateMovieRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 回应的Movie请求
type CreateMovieResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}

// 回应的Err
type ErrMovieResponse struct {
	Status      int    `json:"status"`
	Err         error  `json:"error"`
	Description string `json:"description"`
}

// 回应的而Delete
type DeleteResponse struct {
	DeleteID int   `json:"id"`
	Err      error `json:"error"`
}

// InMem中的Movie,其他由sqlc生成，经仅测试
type MovieDto struct {
	ID    int
	Name  string
	Email string
}

func (s *CreateMovieRequest) ToDto() *MovieDto {
	return &MovieDto{
		ID:    s.ID,
		Name:  s.Name,
		Email: s.Email,
	}
}

func (s *MovieDto) ToResponse(status int, description string) *CreateMovieResponse {
	return &CreateMovieResponse{
		Status:      status,
		Description: description,
		Name:        s.Name,
		Email:       s.Email,
	}
}
