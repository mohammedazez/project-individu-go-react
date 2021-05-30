package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project-individu-go-react/helper"
	"project-individu-go-react/layer/userprofile"
	"strconv"
)

type userProfileHandler struct {
	service userprofile.Service
}

func NewUserProfileHandler(service userprofile.Service) *userProfileHandler {
	return &userProfileHandler{service}
}

func (h *userProfileHandler) GetUserProfileByUserIDHandler(c *gin.Context) {
	userData := int(c.MustGet("currentUser").(int))
	userID := strconv.Itoa(userData)
	userProfile, err := h.service.GetUserProfileByUserID(userID)
	if err != nil {
		responseError := helper.APIResponse("status unauthorize", 401, "error", gin.H{"error": err.Error()})
		c.JSON(401, responseError)
		return
	}

	response := helper.APIResponse("success get user profile by user id", 200, "success", userProfile)
	c.JSON(200, response)
}

func (h *userProfileHandler) SaveNewUserProfileHandler(c *gin.Context) {

	userData := int(c.MustGet("currentUser").(int))

	file, err := c.FormFile("profile") // postman

	if err != nil {
		responseError := helper.APIResponse("status bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	path := fmt.Sprintf("images/profile-%d-%s", userData, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		// log.Println("error line 63")
		responseError := helper.APIResponse("status bad request", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	pathProfileSave := "https://todo-rest-api-golang.herokuapp.com/" + path

	userProfile, err := h.service.SavenewUserProfile(pathProfileSave, userData)

	if err != nil {
		responseError := helper.APIResponse("Internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create user profile", 201, "success", userProfile)
	c.JSON(201, response)
}

func (h *userProfileHandler) UpdateUserProfileByIDHandler(c *gin.Context) {
	userData := int(c.MustGet("currentUser").(int))
	ID := strconv.Itoa(userData)
	//input in postman
	fileImage, err := c.FormFile("profile")
	if err != nil {
		responseError := helper.APIResponse("status bad request", 400, "error", gin.H{"error": err.Error()})
		c.JSON(400, responseError)
		return
	}

	path := fmt.Sprintf("images/profile-%d-%s", userData, fileImage.Filename)
	err = c.SaveUploadedFile(fileImage, path)
	if err != nil {
		responseError := helper.APIResponse("status bad request", 400, "error", gin.H{"error": err.Error()})
		c.JSON(400, responseError)
		return
	}

	pathProfileSave := "https://konsultasi-psikolog.herokuapp.com/" + path
	userProfile, err := h.service.UpdateUserProfileByID(pathProfileSave, ID)
	if err != nil {
		responseError := helper.APIResponse("Internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user profile image", 200, "success update", userProfile)
	c.JSON(200, response)
}
