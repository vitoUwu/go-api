package users

import (
	"errors"
	"net/http"
	"vitooapi/database"
	"vitooapi/models"
	"vitooapi/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func updateUsername(c echo.Context) error {
	db := database.GetConnection()

	oldUsername := c.Param("username")

	user := models.UserModel{}

	query := db.Model(&user).Where("username = ?", oldUsername).First(&user)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.APIError{Message: "User not found"})
		}

		log.Error(query.Error)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occured"})
	}

	requestBody := new(UserBind)

	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid request body"})
	}

	if err := c.Validate(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid request body"})
	}

	if update := query.Update("username", requestBody.Username).Scan(&user); update.Error != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Username is already taken"})
	}

	return c.JSON(http.StatusOK, utils.APIError{Message: "Username updated"})
}
