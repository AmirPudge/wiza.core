package service

import (
	"context"

	"wiza.core/domain"
)

type clientService struct {
	repo domain.ClientRepository
}

func NewClientService(repo domain.ClientRepository) domain.ClientService {
	return &clientService{repo: repo}
}

func (s *clientService) GetByIIN(ctx context.Context, iin string) (*domain.Client, error) {
	return s.repo.GetByIIN(ctx, iin)
}
