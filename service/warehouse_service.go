package service

import (
	"context"
	"errors"
	"stokeria/model/domain"
	"stokeria/model/web"
	"stokeria/repository"

	"github.com/go-playground/validator/v10"
)

type WarehouseService interface {
	Create(ctx context.Context, request web.WarehouseCreateRequest) (web.WarehouseResponse, error)
	Update(ctx context.Context, request web.WarehouseUpdateRequest) (web.WarehouseResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (web.WarehouseResponse, error)
	FindAll(ctx context.Context) ([]web.WarehouseResponse, error)
}

type warehouseServiceImpl struct {
	warehouseRepository repository.WarehouseRepository
	validate            *validator.Validate
}

func NewWarehouseService(warehouseRepository repository.WarehouseRepository, validate *validator.Validate) WarehouseService {
	return &warehouseServiceImpl{
		warehouseRepository: warehouseRepository,
		validate:            validate,
	}
}

func (s *warehouseServiceImpl) Create(ctx context.Context, request web.WarehouseCreateRequest) (web.WarehouseResponse, error) {
	if err := s.validate.Struct(request); err != nil {
		return web.WarehouseResponse{}, err
	}

	existingWarehouse, _ := s.warehouseRepository.FindByCode(ctx, request.Code)
	if existingWarehouse != nil {
		return web.WarehouseResponse{}, errors.New("warehouse code already exists")
	}

	warehouse := domain.Warehouse{
		Code:    request.Code,
		Name:    request.Name,
		Address: request.Address,
	}

	if err := s.warehouseRepository.Create(ctx, &warehouse); err != nil {
		return web.WarehouseResponse{}, err
	}

	return toWarehouseResponse(&warehouse), nil
}

func (s *warehouseServiceImpl) Update(ctx context.Context, request web.WarehouseUpdateRequest) (web.WarehouseResponse, error) {
	if err := s.validate.Struct(request); err != nil {
		return web.WarehouseResponse{}, err
	}

	warehouse, err := s.warehouseRepository.FindById(ctx, request.ID)
	if err != nil {
		return web.WarehouseResponse{}, err
	}

	if warehouse.Code != request.Code {
		existingWarehouse, _ := s.warehouseRepository.FindByCode(ctx, request.Code)
		if existingWarehouse != nil {
			return web.WarehouseResponse{}, errors.New("warehouse code already exists")
		}
	}

	warehouse.Code = request.Code
	warehouse.Name = request.Name
	warehouse.Address = request.Address

	if err := s.warehouseRepository.Update(ctx, warehouse); err != nil {
		return web.WarehouseResponse{}, err
	}

	return toWarehouseResponse(warehouse), nil
}

func (s *warehouseServiceImpl) Delete(ctx context.Context, id string) error {
	return s.warehouseRepository.Delete(ctx, id)
}

func (s *warehouseServiceImpl) FindById(ctx context.Context, id string) (web.WarehouseResponse, error) {
	warehouse, err := s.warehouseRepository.FindById(ctx, id)
	if err != nil {
		return web.WarehouseResponse{}, err
	}
	return toWarehouseResponse(warehouse), nil
}

func (s *warehouseServiceImpl) FindAll(ctx context.Context) ([]web.WarehouseResponse, error) {
	warehouses, err := s.warehouseRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.WarehouseResponse
	for _, warehouse := range warehouses {
		responses = append(responses, toWarehouseResponse(&warehouse))
	}
	return responses, nil
}

func toWarehouseResponse(warehouse *domain.Warehouse) web.WarehouseResponse {
	return web.WarehouseResponse{
		ID:        warehouse.ID,
		Code:      warehouse.Code,
		Name:      warehouse.Name,
		Address:   warehouse.Address,
		CreatedAt: warehouse.CreatedAt,
		UpdatedAt: warehouse.UpdatedAt,
	}
}