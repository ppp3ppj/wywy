package user

import "github.com/labstack/echo/v4"

type UserFrontend struct {
}

func NewUserFrontend(
    g *echo.Group,
) {
    ue := &UserFrontend{}

    g.GET("/signup", ue.Signup)
}
