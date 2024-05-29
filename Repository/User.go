package Repository

import (
	"errors"
	"todo/Model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	BeforeSave(data *Model.User) error
	SaveUser(data *Model.User) error
	GetUserByEmail(email string) (Model.User, error)
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) *userRepository {
	return &userRepository{Db}
}

func ValidatePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *userRepository) BeforeSave(data *Model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	data.Password = string(hashedPassword)
	return nil
}

func (u *userRepository) SaveUser(data *Model.User) error {
	result := u.Db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", data.Username, data.Password, data.Email)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *userRepository) GetUserByEmail(email string) (Model.User, error) {
	user := Model.User{}
	result := u.Db.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&user)
	if result.Error != nil {
		return user, result.Error
	}
	if result.RowsAffected == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}
