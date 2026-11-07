package service

import (
	"context"
	"errors"
	"webook/internal/domain"
	"webook/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("邮箱或密码错误")

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Login(ctx context.Context, email string, password string) (domain.User,error) {
	// 1、先查询是否存在该用户，通过email查询
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{},ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{},err
	}
	// 2、对比密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{},ErrInvalidUserOrPassword
	}
	return u,nil
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}
