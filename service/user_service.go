package service

import (
	"context"
	"errors"
	"stokeria/helper"
	"stokeria/model/domain"
	"stokeria/model/web"
	"stokeria/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error)
	Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (web.UserResponse, error)
	FindAll(ctx context.Context) ([]web.UserResponse, error)
	Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error)
}

type userServiceImpl struct {
	userRepository repository.UserRepository
	validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
		validate:       validate,
	}
}

func (s *userServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) (web.UserResponse, error) {
	if err := s.validate.Struct(request); err != nil {
		return web.UserResponse{}, err
	}

	existingUser, _ := s.userRepository.FindByEmail(ctx, request.Email)
	if existingUser != nil {
		return web.UserResponse{}, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return web.UserResponse{}, err
	}

	user := domain.User{
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
		Name:         request.Name,
		Role:         request.Role,
	}

	if err := s.userRepository.Create(ctx, &user); err != nil {
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(user), nil
}

func (s *userServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error) {
	if err := s.validate.Struct(request); err != nil {
		return web.UserResponse{}, err
	}

	user, err := s.userRepository.FindById(ctx, request.ID)
	if err != nil {
		return web.UserResponse{}, err
	}

	if user.Email != request.Email {
		existingUser, _ := s.userRepository.FindByEmail(ctx, request.Email)
		if existingUser != nil {
			return web.UserResponse{}, errors.New("email already exists")
		}
	}

	user.Email = request.Email
	user.Name = request.Name
	user.Role = request.Role

	if err := s.userRepository.Update(ctx, user); err != nil {
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(*user), nil
}

func (s *userServiceImpl) Delete(ctx context.Context, id string) error {
	return s.userRepository.Delete(ctx, id)
}

func (s *userServiceImpl) FindById(ctx context.Context, id string) (web.UserResponse, error) {
	user, err := s.userRepository.FindById(ctx, id)
	if err != nil {
		return web.UserResponse{}, err
	}
	return helper.ToUserResponse(*user), nil
}

func (s *userServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []web.UserResponse
	for _, user := range users {
		responses = append(responses, helper.ToUserResponse(user))
	}
	return responses, nil
}

func (s *userServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error) {
	if err := s.validate.Struct(request); err != nil {
		return web.UserResponse{}, err
	}

	user, err := s.userRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return web.UserResponse{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return web.UserResponse{}, errors.New("invalid email or password")
	}

	return helper.ToUserResponse(*user), nil
}
