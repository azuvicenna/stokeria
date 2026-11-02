package service

import (
	"context"
	"stokeria/helper"
	"stokeria/model/web"
	"stokeria/repository"
)

type InventoryService interface {
	FindById(ctx context.Context, id string) (web.InventoryResponse, error)
	FindAll(ctx context.Context) ([]web.InventoryResponse, error)
}

type inventoryServiceImpl struct {
	inventoryRepository repository.InventoryRepository
}

func NewInventoryService(inventoryRepository repository.InventoryRepository) InventoryService {
	return &inventoryServiceImpl{
		inventoryRepository: inventoryRepository,
	}
}

func (s *inventoryServiceImpl) FindById(ctx context.Context, id string) (web.InventoryResponse, error) {
	inventory, err := s.inventoryRepository.FindById(ctx, id)
	if err != nil {
		return web.InventoryResponse{}, err
	}
	return helper.ToInventoryResponse(*inventory), nil
}

func (s *inventoryServiceImpl) FindAll(ctx context.Context) ([]web.InventoryResponse, error) {
	inventories, err := s.inventoryRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.InventoryResponse
	for _, inventory := range inventories {
		responses = append(responses, helper.ToInventoryResponse(inventory))
	}
	return responses, nil
}
