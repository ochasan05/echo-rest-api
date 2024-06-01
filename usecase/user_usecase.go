package usecase

import (
	"go-rest-api-todo/model"
	"go-rest-api-todo/repository"
	"go-rest-api-todo/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	LogIn(user model.User) (string, error)
}

type userUsecase struct {
	repository repository.IUserRepository
	validator validator.IUserValidator
}

func NewUserUsecase(repository repository.IUserRepository, validator validator.IUserValidator) IUserUsecase {
	return &userUsecase{repository, validator}
}

func (usecase *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := usecase.validator.UserValidator(user); err != nil {
		return model.UserResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}

	newUser := model.User{Email: user.Email, Password: string(hash)}

	if err := usecase.repository.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse{
		ID: newUser.ID,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (usecase *userUsecase) LogIn(user model.User) (string, error) {
	if err := usecase.validator.UserValidator(user); err != nil {
		return "", err
	}

	storedUser := model.User{}
	if err := usecase.repository.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password),  []byte(user.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}