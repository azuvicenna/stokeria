package service

import (
	"context"
	"stokeria/model/domain"
	"stokeria/model/web"
	"stokeria/repository"

	"github.com/go-playground/validator/v10"
)

type StockAdjustmentService interface {
	Create(ctx context.Context, request web.StockAdjustmentCreateRequest, userID string) (web.StockAdjustmentResponse, error)
	FindById(ctx context.Context, id string) (web.StockAdjustmentResponse, error)
	FindAll(ctx context.Context) ([]web.StockAdjustmentResponse, error)
}

type stockAdjustmentServiceImpl struct {
	stockAdjustmentRepository repository.StockAdjustmentRepository
	validate                  *validator.Validate
}

func NewStockAdjustmentService(stockAdjustmentRepository repository.StockAdjustmentRepository, validate *validator.Validate) StockAdjustmentService {
	return &stockAdjustmentServiceImpl{
		stockAdjustmentRepository: stockAdjustmentRepository,
		validate:                  validate,
	}
}

func (s *stockAdjustmentServiceImpl) Create(ctx context.Context, request web.StockAdjustmentCreateRequest, userID string) (web.StockAdjustmentResponse, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return web.StockAdjustmentResponse{}, err
	}

	adjustment := domain.StockAdjustment{
		ProductID:   request.ProductID,
		WarehouseID: request.WarehouseID,
		UserID:      userID,
		OldQuantity: request.OldQuantity,
		NewQuantity: request.NewQuantity,
		Reason:      request.Reason,
	}

	err = s.stockAdjustmentRepository.Create(ctx, &adjustment)
	if err != nil {
		return web.StockAdjustmentResponse{}, err
	}

	savedAdjustment, err := s.stockAdjustmentRepository.FindById(ctx, adjustment.ID)
	if err != nil {
		return web.StockAdjustmentResponse{}, err
	}

	return toStockAdjustmentResponse(savedAdjustment), nil
}

func (s *stockAdjustmentServiceImpl) FindById(ctx context.Context, id string) (web.StockAdjustmentResponse, error) {
	adjustment, err := s.stockAdjustmentRepository.FindById(ctx, id)
	if err != nil {
		return web.StockAdjustmentResponse{}, err
	}
	return toStockAdjustmentResponse(adjustment), nil
}

func (s *stockAdjustmentServiceImpl) FindAll(ctx context.Context) ([]web.StockAdjustmentResponse, error) {
	adjustments, err := s.stockAdjustmentRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.StockAdjustmentResponse
	for _, adjustment := range adjustments {
		responses = append(responses, toStockAdjustmentResponse(&adjustment))
	}
	return responses, nil
}

func toStockAdjustmentResponse(adjustment *domain.StockAdjustment) web.StockAdjustmentResponse {
	response := web.StockAdjustmentResponse{
		ID:          adjustment.ID,
		OldQuantity: adjustment.OldQuantity,
		NewQuantity: adjustment.NewQuantity,
		Reason:      adjustment.Reason,
		CreatedAt:   adjustment.CreatedAt,
	}

	if adjustment.Product.ID != "" {
		response.Product = &web.ProductResponse{
			ID:   adjustment.Product.ID,
			Name: adjustment.Product.Name,
		}
	}

	if adjustment.Warehouse.ID != "" {
		response.Warehouse = &web.WarehouseResponse{
			ID:   adjustment.Warehouse.ID,
			Name: adjustment.Warehouse.Name,
		}
	}

	if adjustment.User.ID != "" {
		response.User = &web.UserResponse{
			ID:   adjustment.User.ID,
			Name: adjustment.User.Name,
		}
	}

	return response
}