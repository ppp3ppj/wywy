package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/services"
	"github.com/ppp3ppj/wywy/views/auth_views"
)

type AuthService interface {
    CreateUser(user services.User) error
}

func NewAuthHandler(us AuthService) *AuthHandler {
    return &AuthHandler{
        UserService: us,
    }
}

type AuthHandler struct {
    UserService AuthService 
}

func (h *AuthHandler) homeHandler(c echo.Context) error {
    homeView := auth_views.Home()
    return renderView(c, auth_views.HomeIndex(homeView))
}

func (h *AuthHandler) loginHandler(c echo.Context) error {
    loginView := auth_views.Login()
    return renderView(c, auth_views.LoginIndex(loginView))
}

func (h *AuthHandler) registerHandler(c echo.Context) error {
    registerView := auth_views.Register()

    if c.Request().Method == "POST" {
        fmt.Println("POST")
        /*
        user := services.User{
            Email:    c.FormValue("email"),
            Password: c.FormValue("password"),
            Username: c.FormValue("username"),
        }
        */
        user := services.User {
            Email: "test@wywy.com",
            Password: "test3412",
            Username: "test11Agent",
        }
        err := h.UserService.CreateUser(user)
        if err != nil {
            return echo.NewHTTPError(
                echo.ErrInternalServerError.Code,
                fmt.Sprintf("Failed to create user: %v", err),
            )
        }
        return c.Redirect(http.StatusSeeOther, "/login")
    }
    return renderView(c, auth_views.RegisterIndex(registerView))
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
