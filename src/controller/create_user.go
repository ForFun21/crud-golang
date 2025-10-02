package controller

import (
	"github.com/ForFun21/crud-golang/src/configuration/logger"
	"github.com/ForFun21/crud-golang/src/configuration/validation"
	"github.com/ForFun21/crud-golang/src/controller/model/request"
	"github.com/ForFun21/crud-golang/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"))
	var UserRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&UserRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	response := response.UserResponse{
		ID:    "test",
		Email: UserRequest.Email,
		Name:  UserRequest.Name,
		Age:   UserRequest.Age,
	}
	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(200, response)
}
