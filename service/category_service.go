package service

import (
	"context"
	"stokeria/model/domain"
	"stokeria/model/web"
	"stokeria/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error)
	Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (web.CategoryResponse, error)
	FindAll(ctx context.Context) ([]web.CategoryResponse, error)
}

type categoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
	validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &categoryServiceImpl{
		categoryRepository: categoryRepository,
		validate:           validate,
	}
}

func (s *categoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return web.CategoryResponse{}, err
	}

	category := domain.Category{
		Name:        request.Name,
		Description: request.Description,
	}

	err = s.categoryRepository.Create(ctx, &category)
	if err != nil {
		return web.CategoryResponse{}, err
	}

	return toCategoryResponse(&category), nil
}

func (s *categoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return web.CategoryResponse{}, err
	}

	category, err := s.categoryRepository.FindById(ctx, request.ID)
	if err != nil {
		return web.CategoryResponse{}, err
	}

	category.Name = request.Name
	category.Description = request.Description

	err = s.categoryRepository.Update(ctx, category)
	if err != nil {
		return web.CategoryResponse{}, err
	}

	return toCategoryResponse(category), nil
}

func (s *categoryServiceImpl) Delete(ctx context.Context, id string) error {
	return s.categoryRepository.Delete(ctx, id)
}

func (s *categoryServiceImpl) FindById(ctx context.Context, id string) (web.CategoryResponse, error) {
	category, err := s.categoryRepository.FindById(ctx, id)
	if err != nil {
		return web.CategoryResponse{}, err
	}
	return toCategoryResponse(category), nil
}

func (s *categoryServiceImpl) FindAll(ctx context.Context) ([]web.CategoryResponse, error) {
	categories, err := s.categoryRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.CategoryResponse
	for _, category := range categories {
		responses = append(responses, toCategoryResponse(&category))
	}
	return responses, nil
}

func toCategoryResponse(category *domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}