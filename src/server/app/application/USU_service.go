package application

import (
	"github.com/WelintonJunior/SuptecGO/src/server/adapters/domain"
	"github.com/WelintonJunior/SuptecGO/src/server/adapters/repository"
)

type UsuService struct {
	repo repository.UsuRepository
}

func NewUsuService(repo repository.UsuRepository) *UsuService {
	return &UsuService{repo: repo}
}

func (s *UsuService) NewUsu(u domain.USU_users) error {
	return s.repo.NewUsu(u)
}
