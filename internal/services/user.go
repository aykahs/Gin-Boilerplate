package services

import (
	"fmt"

	"github.com/aykahs/Gin-Boilerplate/internal/models"
	"github.com/aykahs/Gin-Boilerplate/internal/pkg/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (UserService *UserService) Me(userID uint) (error, *models.User) {

	me := models.User{
		BasicModel: models.BasicModel{ID: userID},
	}
	res := mysql.GetDB().First(&me, "ID = ?", userID)
	if res.Error != nil || res.RowsAffected == 0 {
		return fmt.Errorf("invalid"), nil
	}
	return nil, &me
}
func (userService *UserService) LoginByUsernamePassword(username string, password string) (string, error) {
	user := models.User{
		Username: username,
	}
	res := mysql.GetDB().First(&user, "username = ?", username)
	if res.Error != nil || res.RowsAffected == 0 {
		return "", res.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("username or password is invalid")
	}
	return "", nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func (userService *UserService) Register(data map[string]interface{}) (*models.User, error) {

	hashedPassword, err := hashPassword(data["password"].(string))
	if err != nil {
		return nil, err
	}
	db := mysql.GetDB()
	var existingUser models.User
	if db.Where("username = ?", data["username"].(string)).First(&existingUser).Error == nil {
		return nil, fmt.Errorf("username already exists")
	}
	user := models.User{
		Name:     data["name"].(string),
		Username: data["username"].(string),
		Password: hashedPassword,
	}
	res := db.Create(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, res.Error
	}

	return &user, nil

}
