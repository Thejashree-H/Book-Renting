package controllers
import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"bookrental/models"
	"bookrental/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usersService *services.UserService
}

func NewUsersController(usersService *services.UserService) *UsersController {
	return &UsersController{usersService: usersService}
}

func (uc UserController) CreateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create User request body", err)
		ctx.AbortWithError(http.StatusInternalError, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Println("Error while unmarshalling create user request body", err)
		ctx.AbortWithError(http.StatusInternalError, err)
		return
	}

	response, err := uc.usersService.CreateUser(&user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (uc UsersController) UpdateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading request body", err)
		ctx.AbortWithError(http.StatusInternalError, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Println("Error while unmarshalling request body", err)
		ctx.AbortWithError(http.StatusInternalError, err)
		return
	}

	response, err := uc.usersService.UpdateUser(&user)
	if repsonseErr != nil {
		ctx.JSON(err.Status, rerr)
		return
	}
	
	ctx.JSON(http.StatusOK, response)
}

func (uc UsersController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := uc.usersService.DeleteUser(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (uc UsersController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := uc.usersService.GetUser(id)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (uc UserController) GetUsersBatch(ctx *gin.Context) {
	response, err := uc.usersService.GetUsersBatch()
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
