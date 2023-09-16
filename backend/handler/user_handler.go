package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"task-app-study/model"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	db *sql.DB
}

type UserForm struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUserHandler(db *sql.DB) *userHandler {
	return &userHandler{db}
}

func (u *userHandler) UserCreate(ctx *gin.Context) {
	var form UserForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, "validate error")
		return
	}
	user := model.NewUser(u.db)
	if user.ExistUserByEmail(form.Email) {
		ctx.JSON(http.StatusBadRequest, "email duplicate")
		return
	}
	user.CreateUser(form.Name, form.Email, form.Password)
}
