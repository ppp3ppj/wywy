package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/views/auth_views"
)

func NewAuthHandler() *AuthHandler {
    return &AuthHandler{}
}

type AuthHandler struct {
}

func (h *AuthHandler) homeHandler(c echo.Context) error {
    homeView := auth_views.Home()
    return renderView(c, auth_views.HomeIndex(homeView))
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
