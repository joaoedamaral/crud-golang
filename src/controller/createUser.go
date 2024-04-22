package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaoedamaral/crud-golang/src/configuration/validation"
	"github.com/joaoedamaral/crud-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	log.Println("[LOGS] Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("[LOGS][CREATEUSER] Error trying to marshal object, error=%s", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)

}
