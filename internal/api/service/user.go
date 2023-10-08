package service

import (
	"github.com/abishz17/go-backend-template/external/auth"
	"github.com/abishz17/go-backend-template/external/password"
	"github.com/abishz17/go-backend-template/infrastructure"
	"github.com/abishz17/go-backend-template/internal/domain"
	"github.com/abishz17/go-backend-template/internal/response"
	"github.com/abishz17/go-backend-template/internal/view"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}

type UserService struct {
	service UserRepository
	env     infrastructure.Env
}

func NewUserService(service UserRepository, env *infrastructure.Env) UserService {
	return UserService{
		service: service,
		env:     *env,
	}
}

func (u *UserService) CreateUser(ctx echo.Context, view view.UserCreateView) (*domain.User, error) {
	hashedPw, err := password.HashPassword(view.Password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Name:     view.UserName,
		Email:    view.Email,
		Password: hashedPw,
	}
	return u.service.Create(user)
}

func (u *UserService) UserLogin(ctx echo.Context, view view.UserLoginView) (*view.LoginResponseView, error) {

	user, err := u.service.GetUserByEmail(view.Email)
	if err != nil {
		err := response.NewAppError("invalid username or password")
		return nil, err
	}
	if err := CheckPassword(view, *user); err != nil {
		return nil, err
	}
	res, err := Login(*user, u.env)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CheckPassword(view view.UserLoginView, user domain.User) error {
	hashedPw := user.Password
	if !password.MatchPassword(hashedPw, view.Password) {
		err := response.NewAppError("invalid username or password")
		return err
	}
	return nil
}

func Login(user domain.User, env infrastructure.Env) (*view.LoginResponseView, error) {
	claims := map[string]interface{}{
		"is_admin": user.IsAdmin,
		"email":    user.Email,
	}
	token, err := auth.GenerateJWTTokens(user, claims, env)
	if err != nil {
		return nil, err
	}
	return &view.LoginResponseView{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil

}
