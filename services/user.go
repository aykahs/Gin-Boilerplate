package services

import (
	"github.com/aykahs/Gin-Boilerplate/models"
	"github.com/aykahs/Gin-Boilerplate/pkg/mysql"
	"github.com/aykahs/Gin-Boilerplate/services/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// @name LoginByUsernamePassword
// @description LoginByUsernamePassword
// @return string
func (userService *UserService) LoginByUsernamePassword(username string, password string) string {

	user := models.User{
		Username: username,
	}
	res := mysql.GetDB().First(&user, "username = ?", username)
	if res.Error != nil || res.RowsAffected == 0 {
		return ""
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return ""
	}
	claims := utils.Claims{
		Username: user.Username,
		Uid:      user.ID,
	}
	token := utils.GenerateToken(&claims)

	return token
}
