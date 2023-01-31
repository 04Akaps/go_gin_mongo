package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_mongo/services"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController {
		UserService : userService
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H("message" : err.Error()))
		return
	}

	err := uc.UserService.CreateUser(&user);

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H("message " : err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Params("name")

	user, err := uc.UserService.GetUser(&username);
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUsers(ctx *gin.Context)  {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context)  {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.UpdateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context)  {
	username := ctx.Param("name")

	err := uc.UserService.DeleteUser(&username)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/getUser/:name", uc.GetUser)
	userRoute.GET("/getAllUsers", uc.GetAllUsers)
	userRoute.DELETE("/deleteUser/:name", uc.DeleteUser)
	userRoute.PUT("/updateUser", uc.UpdateUser)
}