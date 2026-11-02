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

type OutboundTransactionService interface {
	Create(ctx context.Context, request web.OutboundTransactionCreateRequest, userID string) (web.OutboundTransactionResponse, error)
	FindById(ctx context.Context, id string) (web.OutboundTransactionResponse, error)
	FindAll(ctx context.Context) ([]web.OutboundTransactionResponse, error)
}

type outboundTransactionServiceImpl struct {
	outboundRepo repository.OutboundTransactionRepository
	detailRepo   repository.TransactionDetailRepository
	validate     *validator.Validate
}

func NewOutboundTransactionService(
	outboundRepo repository.OutboundTransactionRepository,
	detailRepo repository.TransactionDetailRepository,
	validate *validator.Validate,
) OutboundTransactionService {
	return &outboundTransactionServiceImpl{
		outboundRepo: outboundRepo,
		detailRepo:   detailRepo,
		validate:     validate,
	}
}

func (s *outboundTransactionServiceImpl) Create(ctx context.Context, request web.OutboundTransactionCreateRequest, userID string) (web.OutboundTransactionResponse, error) {
	if err := s.validate.Struct(request); err != nil {
		return web.OutboundTransactionResponse{}, err
	}

	transactionDate, err := time.Parse(time.RFC3339, request.TransactionDate)
	if err != nil {
		return web.OutboundTransactionResponse{}, fmt.Errorf("invalid transaction date format: %v", err)
	}

	transaction := domain.OutboundTransaction{
		ReferenceCode:   fmt.Sprintf("OUT-%d", time.Now().Unix()),
		WarehouseID:     request.WarehouseID,
		Destination:     request.Destination,
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

	if err := s.outboundRepo.Create(ctx, &transaction, details); err != nil {
		return web.OutboundTransactionResponse{}, err
	}

	savedTransaction, err := s.outboundRepo.FindById(ctx, transaction.ID)
	if err != nil {
		return web.OutboundTransactionResponse{}, err
	}

	savedDetails, err := s.detailRepo.FindByTransaction(ctx, transaction.ID, "OUTBOUND")
	if err != nil {
		return web.OutboundTransactionResponse{}, err
	}

	return helper.ToOutboundTransactionResponse(*savedTransaction, savedDetails), nil
}

func (s *outboundTransactionServiceImpl) FindById(ctx context.Context, id string) (web.OutboundTransactionResponse, error) {
	transaction, err := s.outboundRepo.FindById(ctx, id)
	if err != nil {
		return web.OutboundTransactionResponse{}, err
	}

	details, err := s.detailRepo.FindByTransaction(ctx, id, "OUTBOUND")
	if err != nil {
		return web.OutboundTransactionResponse{}, err
	}

	return helper.ToOutboundTransactionResponse(*transaction, details), nil
}

func (s *outboundTransactionServiceImpl) FindAll(ctx context.Context) ([]web.OutboundTransactionResponse, error) {
	transactions, err := s.outboundRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.OutboundTransactionResponse
	for _, transaction := range transactions {
		responses = append(responses, helper.ToOutboundTransactionResponse(transaction, nil))
	}
	return responses, nil
}
