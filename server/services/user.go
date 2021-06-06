package services

import (
	"github.com/CrowderSoup/EchoVue/server/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// UserService interface that user services must implement
type UserService interface {
	Create(*models.User) error
	GetByEmail(string) (*models.User, error)
}

type userService struct {
	db *gorm.DB
}

func newUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (s *userService) Create(user *models.User) error {
	password, err := s.hashPassword(user.Password)
	if err != nil {
		return err
	}

	// Update model with the hashed version of the password before storage
	user.Password = password

	if err = s.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetByEmail gets a user by email
func (s *userService) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (s *userService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword checks if a given password is correct for a user
func (s *userService) checkPassword(password string, user *models.User) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
