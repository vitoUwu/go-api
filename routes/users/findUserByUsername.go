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

func findUserByUsername(c echo.Context) error {
	db := database.GetConnection()
	username := c.Param("username")

	user := models.UserModel{}

	if err := db.Model(&user).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.APIError{Message: "User not found"})
		}

		log.Error(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occured"})
	}

	return c.JSON(http.StatusOK, &user)
}
