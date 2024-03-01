package controler

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"test/model"
	"test/service"

	"github.com/labstack/echo/v4"
)

func AddProfile(c echo.Context) error {
	var addProfile model.Profile
	c.Bind(&addProfile)

	_, err := service.AddProfile(addProfile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.Response{
			Message: "An internal server error",
			Status:  false,
		})
	}
	return c.JSON(http.StatusOK, &model.Response{
		Message: "Successfully added a profile",
		Status:  true,
	})
}

func GetAllProfile(c echo.Context) error {
	profiles, err := service.GetAllProfile()
	if err != nil {
		response := model.Response{
			Message: "An iternal server error, please try again",
			Status:  false,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, profiles)
}

func GetSpecificProfiles(c echo.Context) error {
	idUser, _ := strconv.Atoi(c.Param("id"))
	var userProfile model.Profile

	userProfile, err := service.GetProfile(idUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.Response{
			Message: "An iternal server error occured, please try again",
			Status:  false,
		})

	}
	return c.JSON(http.StatusOK, userProfile)
}

type UserController struct {
	userService *service.UserService
}

type ProfileNotFoundError struct {
	Message string
}

func (e *ProfileNotFoundError) Error() string {
	return e.Message
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func SaveUploadedFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func (c *UserController) UpdatePicture(e echo.Context) error {
	userID, _ := strconv.Atoi(e.Param("id"))
	file, err := e.FormFile("image")
	if err != nil {
		return e.JSON(http.StatusBadRequest, &model.Response{
			Message: "Invalid data ! the data type must be image",
			Status:  false,
		})

	}
	pathImage := "/pictures" + file.Filename
	if err := SaveUploadedFile(file, pathImage); err != nil {
		return e.JSON(http.StatusInternalServerError, &model.Response{
			Message: "An interan server error",
			Status:  false,
		})
	}
	baseURL := "http://localhost:8080"
	pictureURL := baseURL + "/pictures/" + file.Filename

	if err := c.userService.UpdatePictureURL(userID, pictureURL); err != nil {
		return e.JSON(http.StatusInternalServerError, &model.Response{
			Message: "Error uploading the cover image URL",
			Status:  false,
		})
	}
	return e.JSON(http.StatusOK, &model.Response{
		Message: "profile updated successfully",
		Status:  true,
	})
}
