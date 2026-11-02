package service

import (
	"context"
	"stokeria/helper"
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

	return helper.ToStockAdjustmentResponse(*savedAdjustment), nil
}

func (s *stockAdjustmentServiceImpl) FindById(ctx context.Context, id string) (web.StockAdjustmentResponse, error) {
	adjustment, err := s.stockAdjustmentRepository.FindById(ctx, id)
	if err != nil {
		return web.StockAdjustmentResponse{}, err
	}
	return helper.ToStockAdjustmentResponse(*adjustment), nil
}

func (s *stockAdjustmentServiceImpl) FindAll(ctx context.Context) ([]web.StockAdjustmentResponse, error) {
	adjustments, err := s.stockAdjustmentRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.StockAdjustmentResponse
	for _, adjustment := range adjustments {
		responses = append(responses, helper.ToStockAdjustmentResponse(adjustment))
	}
	return responses, nil
}
