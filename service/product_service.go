package service

import (
	"context"
	"stokeria/helper"
	"stokeria/model/domain"
	"stokeria/model/web"
	"stokeria/repository"

	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error)
	Update(ctx context.Context, request web.ProductUpdateRequest) (web.ProductResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (web.ProductResponse, error)
	FindAll(ctx context.Context) ([]web.ProductResponse, error)
}

type productServiceImpl struct {
	productRepository repository.ProductRepository
	validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &productServiceImpl{
		productRepository: productRepository,
		validate:          validate,
	}
}

func (s *productServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return web.ProductResponse{}, err
	}

	product := domain.Product{
		SKU:        request.SKU,
		CategoryID: request.CategoryID,
		Name:       request.Name,
		Price:      request.Price,
	}

	err = s.productRepository.Create(ctx, &product)
	if err != nil {
		return web.ProductResponse{}, err
	}

	savedProduct, err := s.productRepository.FindById(ctx, product.ID)
	if err != nil {
		return web.ProductResponse{}, err
	}

	return helper.ToProductResponse(*savedProduct), nil
}

func (s *productServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) (web.ProductResponse, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return web.ProductResponse{}, err
	}

	product, err := s.productRepository.FindById(ctx, request.ID)
	if err != nil {
		return web.ProductResponse{}, err
	}

	product.SKU = request.SKU
	product.CategoryID = request.CategoryID
	product.Name = request.Name
	product.Price = request.Price

	err = s.productRepository.Update(ctx, product)
	if err != nil {
		return web.ProductResponse{}, err
	}

	updatedProduct, err := s.productRepository.FindById(ctx, product.ID)
	if err != nil {
		return web.ProductResponse{}, err
	}

	return helper.ToProductResponse(*updatedProduct), nil
}

func (s *productServiceImpl) Delete(ctx context.Context, id string) error {
	return s.productRepository.Delete(ctx, id)
}

func (s *productServiceImpl) FindById(ctx context.Context, id string) (web.ProductResponse, error) {
	product, err := s.productRepository.FindById(ctx, id)
	if err != nil {
		return web.ProductResponse{}, err
	}
	return helper.ToProductResponse(*product), nil
}

func (s *productServiceImpl) FindAll(ctx context.Context) ([]web.ProductResponse, error) {
	products, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.ToProductResponses(products), nil
}
