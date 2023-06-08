package users

import (
	"errors"
	"log"
	"net/http"
	"vitooapi/database"
	"vitooapi/models"
	"vitooapi/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func createUser(c echo.Context) error {
	requestBody := new(UserBind)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid request body"})
	}

	if err := c.Validate(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid request body"})
	}

	user := models.UserModel{}

	db := database.GetConnection()

	err := db.Model(&models.UserModel{}).Where("username = ?", &requestBody.Username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = models.UserModel{
			Username: requestBody.Username,
		}

		createUserErr := db.Create(&user).Error
		if createUserErr != nil {
			log.Fatal(createUserErr)
		}

		return c.JSON(http.StatusCreated, user)
	} else {
		return c.JSON(http.StatusConflict, utils.APIError{Message: "This username is already taken"})
	}
}
