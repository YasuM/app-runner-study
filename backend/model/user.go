package model

import (
	"context"
	"database/sql"
	"log"
	"task-app-study/entity"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *user {
	return &user{db: db}
}

func (u *user) ExistUserByEmail(email string) bool {
	query := entity.New(u.db)
	cnt, _ := query.CountUserByEmail(context.Background(), email)
	return cnt > 0
}

func (u *user) CreateUser(name, email, password string) error {
	query := entity.New(u.db)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	params := entity.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: string(hashPassword),
	}
	_, err = query.CreateUser(context.Background(), params)
	return err
}
