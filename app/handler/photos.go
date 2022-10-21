package handler

import (
	"final_assignment/app/helpers"
	"final_assignment/app/models"
	"final_assignment/app/repository"
	"final_assignment/app/resource"
	"final_assignment/app/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PhotoHandler struct {
	repo repository.PhotoRepository
}

func NewPhotoHandler() *PhotoHandler {
	return &PhotoHandler{
		repository.NewPhotoRepository(),
	}
}

type PhotoOut struct {
	ID           uint      `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	PhotoAt      time.Time `json:"ordered_at"`
	Items        []ItemOut `gorm:"foreignKey:PhotoID"`
}

func (h *PhotoHandler) NewPhoto(c *gin.Context) {
	repo := h.repo
	var req resource.NewPhoto
	err := c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("Error Request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.GetInt("UserID")
	var Photo models.Photo
	Photo.UserID = uint(userId)
	err = repo.NewPhoto(&Photo, req)
	if err != nil {
		response := helpers.APIResponse2("Failed", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Successed", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
	c.JSON(http.StatusOK, response)
}

func (h *PhotoHandler) GetPhotos(c *gin.Context) {
	repo := h.repo
	userId := c.GetInt("UserID")
	var Photos []models.Photo
	err := repo.GetPhotos(&Photos, uint(userId))
	if err != nil {
		response := helpers.APIResponse2("Failed", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	var photoList []map[string]interface{}
	for _, eachPhoto := range Photos {
		data := map[string]interface{}{
			"id":         eachPhoto.ID,
			"title":      eachPhoto.Title,
			"caption":    eachPhoto.Caption,
			"photo_url":  eachPhoto.PhotoUrl,
			"user_id":    eachPhoto.UserID,
			"created_at": eachPhoto.CreatedAt,
		}
		photoList = append(photoList, data)
	}
	response := helpers.APIResponse2("Successed", http.StatusOK, 0, 0, 0, photoList)
	c.JSON(http.StatusOK, response)
}

func (h *PhotoHandler) RemovePhoto(c *gin.Context) {
	photoId := c.Param("photoId")
	photoIdInt, err := strconv.Atoi(photoId)
	repo := h.repo
	userId := c.GetInt("UserID")
	err = repo.RemovePhoto(uint(userId), uint(photoIdInt))
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, err)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Remove Successful.", http.StatusOK, 0, 0, 0, "")
	c.JSON(http.StatusOK, response)
}

func (h *PhotoHandler) EditPhoto(c *gin.Context) {
	photoId := c.Param("photoId")
	photoIdInt, err := strconv.Atoi(photoId)
	repo := h.repo
	var req resource.NewPhoto
	err = c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.GetInt("UserID")
	var Photo models.Photo
	Photo.ID = uint(photoIdInt)
	err = repo.EditPhoto(&Photo, req, uint(userId))
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Successed", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.UpdatedAt,
	})
	c.JSON(http.StatusOK, response)
}
