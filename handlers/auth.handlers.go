package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/services"
	"github.com/ppp3ppj/wywy/views/auth_views"
	"github.com/ppp3ppj/wywy/views/auth_views/register_components"
	"golang.org/x/crypto/bcrypt"
)

const (
    auth_sessios_key string = "authenticate-sessions"
    auth_key string = "authenticated"
    user_id_key string = "user_id"
    username_key string = "username"
    tz_key string = "time_zone"
)

type AuthService interface {
    CreateUser(user services.User) error
    ValidateEmail(email string) error
    CheckEmail(email string) (services.User, error)
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

    if c.Request().Method == "POST" {
        // obtaining the time zone from POST Request
        tzone := ""
        if len(c.Request().Header["X-Timezone"]) != 0 {
            tzone = c.Request().Header["X-Timezone"][0]
        }
        email := c.FormValue("email")
        //password := c.FormValue("password")
        user, err := h.UserService.CheckEmail(email)
        fmt.Println("user is ", user)
        if err != nil {
            if strings.Contains(err.Error(), "email not found") {
                return c.Redirect(http.StatusSeeOther, "/login")
            }
            return echo.NewHTTPError(
                echo.ErrInternalServerError.Code,
                fmt.Sprintf("Failed to check email: %v", err),
            )
        }

        err = bcrypt.CompareHashAndPassword(
            []byte(user.Password), 
            []byte(c.FormValue("password")),
        )

        if err != nil {
            return c.Redirect(http.StatusSeeOther, "/login")
        }

        // Get Session and seitting Cookies
        sess, _ := session.Get(auth_sessios_key, c)
        sess.Options = &sessions.Options{
            Path: "/",
            MaxAge: 3600, // in seconds
            HttpOnly: true,
        }

        // Set user as authenticated

        fmt.Println("tz_key is ", tzone)
        sess.Values = map[interface{}]interface{}{
            auth_key: true,
            user_id_key: user.Id,
            username_key: user.Username,
            tz_key: tzone,
        }
        sess.Save(c.Request(), c.Response())

        return c.Redirect(http.StatusSeeOther, "/")
    }
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
            Email: c.FormValue("email"),
            Password: c.FormValue("password"),
            Username: c.FormValue("username"),
            Firstname: c.FormValue("firstname"),
            Lastname: c.FormValue("lastname"),
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
    email := c.FormValue("email")
    if c.Request().Method == "POST" {
        if len(strings.TrimSpace(email)) == 0 {
            return renderView(c, register_components.EmailInlineValidation(false ,email))
        }
        result := h.UserService.ValidateEmail(email)
        if strings.Contains(result.Error(), "email found") {
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
