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

type SocialMediaHandler struct {
	repo repository.SocialMediaRepository
}

func NewSocialMediaHandler() *SocialMediaHandler {
	return &SocialMediaHandler{
		repository.NewSocialMediaRepository(),
	}
}

type SocialMediaOut struct {
	ID              uint      `json:"order_id"`
	CustomerName    string    `json:"customer_name"`
	SocialMediaedAt time.Time `json:"ordered_at"`
	Items           []ItemOut `gorm:"foreignKey:SocialMediaID"`
}

func (h *SocialMediaHandler) NewSocialMedia(c *gin.Context) {
	repo := h.repo
	var req resource.NewSocialMedia
	err := c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.GetInt("UserID")
	var SocialMedia models.SocialMedia
	SocialMedia.UserID = uint(userId)
	err = repo.NewSocialMedia(&SocialMedia, req)
	if err != nil {
		response := helpers.APIResponse2("Failed", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Successed", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":                SocialMedia.ID,
		"name":              SocialMedia.Name,
		"social_medial_url": SocialMedia.SocialMedialUrl,
		"user_id":           SocialMedia.UserID,
		"created_at":        SocialMedia.CreatedAt,
	})
	c.JSON(http.StatusOK, response)
}

func (h *SocialMediaHandler) GetSocialMedia(c *gin.Context) {
	repo := h.repo
	userId := c.GetInt("UserID")
	var SocialMedias []models.SocialMedia
	err := repo.GetSocialMedia(&SocialMedias, uint(userId))
	if err != nil {
		response := helpers.APIResponse2("Failed", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	var photoList []map[string]interface{}
	for _, eachSocialMedia := range SocialMedias {
		data := map[string]interface{}{
			"id":               eachSocialMedia.ID,
			"name":             eachSocialMedia.Name,
			"social_media_url": eachSocialMedia.SocialMedialUrl,
			"user_id":          eachSocialMedia.UserID,
			"created_at":       eachSocialMedia.CreatedAt,
			"updated_at":       eachSocialMedia.UpdatedAt,
			"user": map[string]interface{}{
				"id":       eachSocialMedia.User.ID,
				"email":    eachSocialMedia.User.Email,
				"username": eachSocialMedia.User.Username,
			},
		}
		photoList = append(photoList, data)
	}
	response := helpers.APIResponse2("Successed", http.StatusOK, 0, 0, 0, photoList)
	c.JSON(http.StatusOK, response)
}

func (h *SocialMediaHandler) RemoveSocialMedia(c *gin.Context) {
	socialMediaId := c.Param("socialMediaId")
	socialMediaIdInt, err := strconv.Atoi(socialMediaId)
	repo := h.repo
	userId := c.GetInt("UserID")
	err = repo.RemoveSocialMedia(uint(userId), uint(socialMediaIdInt))
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, err)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Remove Successful.", http.StatusOK, 0, 0, 0, "")
	c.JSON(http.StatusOK, response)
}

func (h *SocialMediaHandler) EditSocialMedia(c *gin.Context) {
	socialMediaId := c.Param("socialMediaId")
	socialMediaIdInt, err := strconv.Atoi(socialMediaId)
	repo := h.repo
	var req resource.EditSocialMedia
	err = c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.GetInt("UserID")
	var SocialMedia models.SocialMedia
	SocialMedia.ID = uint(socialMediaIdInt)
	err = repo.EditSocialMedia(&SocialMedia, req, uint(userId))
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Edit Social Success", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":                SocialMedia.ID,
		"name":              SocialMedia.Name,
		"social_medial_url": SocialMedia.SocialMedialUrl,
		"user_id":           SocialMedia.UserID,
		"updated_at":        SocialMedia.UpdatedAt,
	})
	c.JSON(http.StatusOK, response)
}
