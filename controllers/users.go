package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indahouse-aleksandr/mp_backend/inout"
	"github.com/indahouse-aleksandr/mp_backend/repositories"
	"github.com/indahouse-aleksandr/mp_backend/services"
)

func GetUserByID(c *gin.Context) {

	var inputModel inout.InputGetUserById
	if err := c.BindUri(&inputModel); err != nil {
		c.JSON(http.StatusBadRequest, inout.ErrMsgValidator(err))
		return
	}

	c.JSON(http.StatusAccepted, &inputModel)
	return
}

func CreateUser(c *gin.Context) {

	var inputModel inout.InputCreateUser
	if err := c.ShouldBindJSON(&inputModel); err != nil {
		c.JSON(http.StatusBadRequest, inout.ErrMsgValidator(err))
		return
	}

	r := repositories.NewPgx()
	srv := services.NewUserService(c, r)
	row, err := srv.CreateUser(inputModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, inout.ErrMsgUser(err))
		return
	}

	c.JSON(http.StatusBadRequest, inout.CreateUserOk(row))
}
