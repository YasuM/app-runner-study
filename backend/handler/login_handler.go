package handler

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"task-app-study/entity"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginHandler struct {
	db *sql.DB
}

type LoginForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewLoginHandler(db *sql.DB) *loginHandler {
	return &loginHandler{db}
}

func (l *loginHandler) Login(ctx *gin.Context) {
	var form LoginForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, "validate error")
		return
	}
	query := entity.New(l.db)

	password, err := query.GetUserPasswordByEmail(context.Background(), form.Email)
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(form.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "login error")
		return
	}
	ctx.JSON(http.StatusOK, "login success")
}
