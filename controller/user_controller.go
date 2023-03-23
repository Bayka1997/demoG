package controller

import (
	"demoG/data/request"
	"demoG/data/response"
	"demoG/helper"
	"demoG/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{
		usersService: service,
	}
}

// CreateUsers		godoc
// @Summary			Create users
// @Description		Save users data in Db.
// @Param			users body request.CreateUsersRequest true "Create users"
// @Produce			application/json
// @Users			users
// @Success			200 {object} response.Response{}
// @Router			/users [post]
func (controller *UsersController) Create(ctx *gin.Context) {
	log.Info().Msg("create users")
	createUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.ErrorPanic(err)

	controller.usersService.Create(createUsersRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateUsers		godoc
// @Summary			Update users
// @Description		Update users data.
// @Param			userId path string true "update users by id"
// @Param			users body request.CreateUsersRequest true  "Update users"
// @Users			users
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/users/{userId} [patch]
func (controller *UsersController) Update(ctx *gin.Context) {
	log.Info().Msg("update users")
	updateUsersRequest := request.UpdateUsersRequest{}
	err := ctx.ShouldBindJSON(&updateUsersRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	updateUsersRequest.Id = id

	controller.usersService.Update(updateUsersRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteUsers		godoc
// @Summary			Delete users
// @Description		Remove users data by id.
// @Produce			application/json
// @Users			users
// @Success			200 {object} response.Response{}
// @Router			/users/{userID} [delete]
func (controller *UsersController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete users")
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.usersService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdUsers 		godoc
// @Summary				Get Single users by id.
// @Param				userId path string true "update users by id"
// @Description			Return the tahs whoes userId valu mathes id.
// @Produce				application/json
// @Users				users
// @Success				200 {object} response.Response{}
// @Router				/users/{userId} [get]
func (controller *UsersController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid users")
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.usersService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllUsers 		godoc
// @Summary			Get All users.
// @Description		Return list of users.
// @Users			users
// @Success			200 {obejct} response.Response{}
// @Router			/users [get]
func (controller *UsersController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll users")
	userResponse := controller.usersService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
