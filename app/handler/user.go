package handler

import (
	"final_assignment/app/helpers"
	"final_assignment/app/models"
	"final_assignment/app/repository"
	"final_assignment/app/resource"
	"final_assignment/app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		repository.NewUserRepository(),
	}
}

type UserOut struct {
	ID           uint      `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	UseredAt     time.Time `json:"ordered_at"`
	Items        []ItemOut `gorm:"foreignKey:UserID"`
}

type ItemOut struct {
	ItemID      uint   `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	UserID      uint   `json:"order_id"`
}

func (h *UserHandler) Signup(c *gin.Context) {
	repo := h.repo
	var req resource.NewUser
	err := c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var User models.User
	err = repo.Signup(&User, req)
	if err != nil {
		response := helpers.APIResponse2("This email is already registered", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Successed", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Signin(c *gin.Context) {
	repo := h.repo
	var req resource.Login
	err := c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token, err := repo.Signin(req.Email, req.Password)
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, err)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Signin Successful", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"token": token,
	})
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) RemoveUser(c *gin.Context) {
	repo := h.repo
	userId := c.GetInt("UserID")
	err := repo.RemoveUser(userId)
	if err != nil {
		response := helpers.APIResponse2(err.Error(), http.StatusBadRequest, 0, 0, 0, err)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Delete Successful", http.StatusOK, 0, 0, 0, "")
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) EditUser(c *gin.Context) {
	repo := h.repo
	var req resource.EditUser
	err := c.ShouldBind(&req)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error_messages": errors}
		response := helpers.APIResponse("bad request", http.StatusBadRequest, "error", errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := c.GetInt("UserID")
	var User models.User
	User.ID = uint(userId)
	err = repo.EditUser(&User, req)
	if err != nil {
		response := helpers.APIResponse2("Failed", http.StatusBadRequest, 0, 0, 0, "")
		c.JSON(http.StatusOK, response)
		return
	}
	response := helpers.APIResponse2("Successed", http.StatusOK, 0, 0, 0, map[string]interface{}{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
	})
	c.JSON(http.StatusOK, response)
}
