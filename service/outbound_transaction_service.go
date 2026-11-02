package service

import (
	"context"
	"fmt"
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

	return toOutboundTransactionResponse(savedTransaction, savedDetails), nil
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

	return toOutboundTransactionResponse(transaction, details), nil
}

func (s *outboundTransactionServiceImpl) FindAll(ctx context.Context) ([]web.OutboundTransactionResponse, error) {
	transactions, err := s.outboundRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.OutboundTransactionResponse
	for _, transaction := range transactions {
		responses = append(responses, toOutboundTransactionResponse(&transaction, nil))
	}
	return responses, nil
}

func toOutboundTransactionResponse(transaction *domain.OutboundTransaction, details []domain.TransactionDetail) web.OutboundTransactionResponse {
	response := web.OutboundTransactionResponse{
		ID:              transaction.ID,
		ReferenceCode:   transaction.ReferenceCode,
		Destination:     transaction.Destination,
		TransactionDate: transaction.TransactionDate,
		Status:          transaction.Status,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}

	if transaction.User.ID != "" {
		response.User = &web.UserResponse{
			ID:   transaction.User.ID,
			Name: transaction.User.Name,
		}
	}

	if len(details) > 0 {
		var detailResponses []web.TransactionDetailResponse
		for _, d := range details {
			detailRes := web.TransactionDetailResponse{
				ID:        d.ID,
				Quantity:  d.Quantity,
				UnitPrice: d.UnitPrice,
			}
			if d.Product.ID != "" {
				detailRes.Product = &web.ProductResponse{
					ID:   d.Product.ID,
					Name: d.Product.Name,
				}
			}
			detailResponses = append(detailResponses, detailRes)
		}
		response.Details = detailResponses
	}

	return response
}