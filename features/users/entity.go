package users

import "github.com/labstack/echo/v4"

type Core struct {
	ID   uint
	Nama string
}

type UserHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type UserLogic interface {
	AddUser(newUser Core) (Core, error)
	Update()
	Delete()
}

type UserData interface {
	Insert(newUser Core) (Core, error)
	Update(updatedData Core)
	Delete()
}
