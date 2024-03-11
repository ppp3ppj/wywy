package handlers

import (
	//"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var (
    fromProtected bool = false
    isError bool = false
)

func SetupRoutes(e *echo.Echo, h *AuthHandler) {
    e.GET("/", h.homeHandler)
    e.GET("/login", h.loginHandler)
    e.GET("/register", h.registerHandler)
    e.POST("/register", h.registerHandler)
    e.POST("/login", h.loginHandler)

    registerGroup := e.Group("/register")
    registerGroup.POST("/email", h.registerEmailHandler)

}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
