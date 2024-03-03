package handlers

import (
	//"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
//	"github.com/ppp3ppj/wywy/views/auth_views"
)

var (
    fromProtected bool = false
    isError bool = false
)

func SetupRoutes(e *echo.Echo) {
    //component := auth_views.HomeIndex()
    //e.GET("/", Render(e.AcquireContext().Request(), http.StatusOK, component))
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
