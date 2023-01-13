package users

import "github.com/labstack/echo/v4"

type Core struct {
	ID   uint
	Nama string
}

type UserHandler interface {
	Add() echo.HandlerFunc
}

type UserLogic interface {
	AddUser(newUser Core) (Core, error)
}

type UserData interface {
	Insert(newUser Core) (Core, error)
}
