package controller

import (
	"net/http"

	"github.com/ForFun21/crud-golang/src/configuration/logger"
	"github.com/ForFun21/crud-golang/src/configuration/validation"
	"github.com/ForFun21/crud-golang/src/controller/model/request"
	"github.com/ForFun21/crud-golang/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(
		UserRequest.Email,
		UserRequest.Password,
		UserRequest.Name,
		UserRequest.Age,
	)
	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
