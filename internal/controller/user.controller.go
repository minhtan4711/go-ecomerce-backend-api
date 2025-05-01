package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecomerce-backend-api/internal/service"
	"github.com/go-ecomerce-backend-api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> model -> db
func (uc *UserController) GetUserByID(c *gin.Context) {
	// if err != nil {
	// 	response.ErrorResponse(c, 20003, "no need")
	// }
	response.SuccessResponse(c, 20001, "success")
}
