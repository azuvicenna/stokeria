package service

import (
	"context"
	"stokeria/model/domain"
	"stokeria/model/web"
	"stokeria/repository"

	"github.com/go-playground/validator/v10"
)

type SupplierService interface {
	Create(ctx context.Context, request web.SupplierCreateRequest) (web.SupplierResponse, error)
	Update(ctx context.Context, request web.SupplierUpdateRequest) (web.SupplierResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (web.SupplierResponse, error)
	FindAll(ctx context.Context) ([]web.SupplierResponse, error)
}

type supplierServiceImpl struct {
	supplierRepository repository.SupplierRepository
	validate           *validator.Validate
}

func NewSupplierService(supplierRepository repository.SupplierRepository, validate *validator.Validate) SupplierService {
	return &supplierServiceImpl{
		supplierRepository: supplierRepository,
		validate:           validate,
	}
}

func (s *supplierServiceImpl) Create(ctx context.Context, request web.SupplierCreateRequest) (web.SupplierResponse, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return web.SupplierResponse{}, err
	}

	supplier := domain.Supplier{
		Name:          request.Name,
		ContactPerson: request.ContactPerson,
		Phone:         request.Phone,
		Address:       request.Address,
	}

	err = s.supplierRepository.Create(ctx, &supplier)
	if err != nil {
		return web.SupplierResponse{}, err
	}

	return toSupplierResponse(&supplier), nil
}

func (s *supplierServiceImpl) Update(ctx context.Context, request web.SupplierUpdateRequest) (web.SupplierResponse, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return web.SupplierResponse{}, err
	}

	supplier, err := s.supplierRepository.FindById(ctx, request.ID)
	if err != nil {
		return web.SupplierResponse{}, err
	}

	supplier.Name = request.Name
	supplier.ContactPerson = request.ContactPerson
	supplier.Phone = request.Phone
	supplier.Address = request.Address

	err = s.supplierRepository.Update(ctx, supplier)
	if err != nil {
		return web.SupplierResponse{}, err
	}

	return toSupplierResponse(supplier), nil
}

func (s *supplierServiceImpl) Delete(ctx context.Context, id string) error {
	return s.supplierRepository.Delete(ctx, id)
}

func (s *supplierServiceImpl) FindById(ctx context.Context, id string) (web.SupplierResponse, error) {
	supplier, err := s.supplierRepository.FindById(ctx, id)
	if err != nil {
		return web.SupplierResponse{}, err
	}
	return toSupplierResponse(supplier), nil
}

func (s *supplierServiceImpl) FindAll(ctx context.Context) ([]web.SupplierResponse, error) {
	suppliers, err := s.supplierRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.SupplierResponse
	for _, supplier := range suppliers {
		responses = append(responses, toSupplierResponse(&supplier))
	}
	return responses, nil
}

func toSupplierResponse(supplier *domain.Supplier) web.SupplierResponse {
	return web.SupplierResponse{
		ID:            supplier.ID,
		Name:          supplier.Name,
		ContactPerson: supplier.ContactPerson,
		Phone:         supplier.Phone,
		Address:       supplier.Address,
		CreatedAt:     supplier.CreatedAt,
		UpdatedAt:     supplier.UpdatedAt,
	}
}