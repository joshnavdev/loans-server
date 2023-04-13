package service

import (
	"errors"
	"time"

	"github.com/joshnavdev/loans-server/models"
	repository "github.com/joshnavdev/loans-server/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
  userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
  return AuthService{
    userRepository: userRepository,
  }
}

func (service *AuthService) SignUp(user models.User) error {
  userCount, err := service.userRepository.CountByEmail(user.Email)
  if err != nil {
    return err // Email is been used 400
  }

  if userCount > 0 {
    return errors.New("Email is in used")
  }

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
    return err // Something is wrong 500
  }

  user.Password = string(hashedPassword)
  user.Created = time.Now()

  if err := service.userRepository.Create(user); err != nil {
    return err // Something is wrong 500
  }

  return nil
}
