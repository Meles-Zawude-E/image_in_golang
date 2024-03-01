package routes

import (
	"test/controler"
	"test/database"
	"test/service"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()
	db := database.Connection()

	userService := service.NewUserService(db)
	userController := controler.NewUserController(userService)

	//define api endpoints
	e.POST("/profile/add", controler.AddProfile)
	e.GET("/profiles/", controler.GetAllProfile)
	e.GET("/profile/:id", controler.GetSpecificProfiles)
	e.POST("/profile/:id/upload", userController.UpdatePicture)

	e.Static("/picture", "/pictures")

	return e
}
