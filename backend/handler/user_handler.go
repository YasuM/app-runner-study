package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"task-app-study/entity"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	query := entity.New(u.db)
	cnt, _ := query.CountUserByEmail(context.Background(), form.Email)
	if cnt > 0 {
		ctx.JSON(http.StatusBadRequest, "email duplicate")
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	params := entity.CreateUserParams{
		Name:     form.Name,
		Email:    form.Email,
		Password: string(hashPassword),
	}
	_, err = query.CreateUser(context.Background(), params)
}
