package service

import (
	"context"
	"fmt"
	"stokeria/helper"
	"stokeria/model/domain"
	"stokeria/model/web"
	"stokeria/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type InboundTransactionService interface {
	Create(ctx context.Context, request web.InboundTransactionCreateRequest, userID string) (web.InboundTransactionResponse, error)
	FindById(ctx context.Context, id string) (web.InboundTransactionResponse, error)
	FindAll(ctx context.Context) ([]web.InboundTransactionResponse, error)
}

type inboundTransactionServiceImpl struct {
	inboundRepo repository.InboundTransactionRepository
	detailRepo  repository.TransactionDetailRepository
	validate    *validator.Validate
}

func NewInboundTransactionService(
	inboundRepo repository.InboundTransactionRepository,
	detailRepo repository.TransactionDetailRepository,
	validate *validator.Validate,
) InboundTransactionService {
	return &inboundTransactionServiceImpl{
		inboundRepo: inboundRepo,
		detailRepo:  detailRepo,
		validate:    validate,
	}
}

func (s *inboundTransactionServiceImpl) Create(ctx context.Context, request web.InboundTransactionCreateRequest, userID string) (web.InboundTransactionResponse, error) {
	if err := s.validate.Struct(request); err != nil {
		return web.InboundTransactionResponse{}, err
	}

	transactionDate, err := time.Parse(time.RFC3339, request.TransactionDate)
	if err != nil {
		return web.InboundTransactionResponse{}, fmt.Errorf("invalid transaction date format: %v", err)
	}

	transaction := domain.InboundTransaction{
		ReferenceCode:   fmt.Sprintf("INB-%d", time.Now().Unix()),
		SupplierID:      request.SupplierID,
		WarehouseID:     request.WarehouseID,
		UserID:          userID,
		TransactionDate: transactionDate,
		Status:          "COMPLETED",
	}

	var details []domain.TransactionDetail
	for _, reqDetail := range request.Details {
		details = append(details, domain.TransactionDetail{
			ProductID: reqDetail.ProductID,
			Quantity:  reqDetail.Quantity,
			UnitPrice: reqDetail.UnitPrice,
		})
	}

	if err := s.inboundRepo.Create(ctx, &transaction, details); err != nil {
		return web.InboundTransactionResponse{}, err
	}

	savedTransaction, err := s.inboundRepo.FindById(ctx, transaction.ID)
	if err != nil {
		return web.InboundTransactionResponse{}, err
	}

	savedDetails, err := s.detailRepo.FindByTransaction(ctx, transaction.ID, "INBOUND")
	if err != nil {
		return web.InboundTransactionResponse{}, err
	}

	return helper.ToInboundTransactionResponse(*savedTransaction, savedDetails), nil
}

func (s *inboundTransactionServiceImpl) FindById(ctx context.Context, id string) (web.InboundTransactionResponse, error) {
	transaction, err := s.inboundRepo.FindById(ctx, id)
	if err != nil {
		return web.InboundTransactionResponse{}, err
	}

	details, err := s.detailRepo.FindByTransaction(ctx, id, "INBOUND")
	if err != nil {
		return web.InboundTransactionResponse{}, err
	}

	return helper.ToInboundTransactionResponse(*transaction, details), nil
}

func (s *inboundTransactionServiceImpl) FindAll(ctx context.Context) ([]web.InboundTransactionResponse, error) {
	transactions, err := s.inboundRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.InboundTransactionResponse
	for _, transaction := range transactions {
		responses = append(responses, helper.ToInboundTransactionResponse(transaction, nil))
	}
	return responses, nil
}
