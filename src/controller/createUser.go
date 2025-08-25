package controller

import (
	"fmt"
	"log"

	"github.com/ForFun21/crud-golang/src/configuration/rest_err/validation"
	"github.com/ForFun21/crud-golang/src/controller/model/request"
	"github.com/ForFun21/crud-golang/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controller")
	var UserRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&UserRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	fmt.Println(UserRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: UserRequest.Email,
		Name:  UserRequest.Name,
		Age:   UserRequest.Age,
	}
	c.JSON(200, response)
}
