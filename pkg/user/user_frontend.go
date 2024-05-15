package user

import (
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/pkg/models"
)

type UserFrontend struct {
    userRepo models.UserRepository
}

func NewUserFrontend(
    g *echo.Group,
    userRepo models.UserRepository,
) {
    ue := &UserFrontend{
        userRepo: userRepo,
    }

    g.GET("/signup", ue.Signup)
    g.POST("/signup", ue.Signup)
}
