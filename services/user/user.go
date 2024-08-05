package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/o-ga09/note-app-backendapi/domain"
)

type UserService struct {
	userRepo domain.IUserRepository
}

func NewUserService(userRepo domain.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) FetchUserById(ctx context.Context, id string) (*domain.User, error) {
	res, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) FetchUsers(ctx context.Context) ([]*domain.User, error) {
	res, err := s.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*domain.User, 0)
	for _, user := range res {
		user := domain.User{
			UserID:    user.UserID,
			Username:  user.Username,
			UserEmail: user.UserEmail,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		users = append(users, &user)
	}
	return users, nil
}

func (s *UserService) CreateUser(ctx context.Context, username, userEmail string) error {
	user := domain.User{
		Username:  username,
		UserEmail: userEmail,
	}
	err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uuid.UUID, username, userEmail string) error {
	user := domain.User{
		UserID:    id,
		Username:  username,
		UserEmail: userEmail,
	}
	err := s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUserById(ctx context.Context, id string) error {
	err := s.userRepo.DeleteUserById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
