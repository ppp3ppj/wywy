package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/services"
	"github.com/ppp3ppj/wywy/views/auth_views"
	"github.com/ppp3ppj/wywy/views/auth_views/register_components"
)

type AuthService interface {
    CreateUser(user services.User) error
    ValidateEmail(email string) error
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
            if strings.Contains(err.Error(), "username has been used") {
                err = errors.New("Username has been used")

                return c.Redirect(http.StatusSeeOther, "/register")
            }
            /*
            return echo.NewHTTPError(
                echo.ErrInternalServerError.Code,
                fmt.Sprintf("Failed to create user: %v", err),
            )
            */
        }
        return c.Redirect(http.StatusSeeOther, "/login")
    }
    return renderView(c, auth_views.RegisterIndex(registerView))
}

func (h *AuthHandler) registerEmailHandler(c echo.Context) error {


    if c.Request().Method == "POST" {

    email := c.FormValue("email")
    fmt.Println("POST RegisterEmailHandler ppp-agent", email)
    if len(strings.TrimSpace(email)) == 0 {
        fmt.Println("email is empty by ppppppp")
        return renderView(c, register_components.EmailInlineValidation(false ,email))
    }
        fmt.Println("POST RegisterEmailHandler", email)
        result := h.UserService.ValidateEmail(email)
        fmt.Println("result is ", result)
        if strings.Contains(result.Error(), "email found") {
            fmt.Println("email found", email)
            return renderView(c, register_components.EmailInlineValidation(false ,email))
        }
        return renderView(c, register_components.EmailInlineValidation(true, email))
    }
    return nil
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
