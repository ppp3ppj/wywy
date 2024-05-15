package user

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/pkg/models"
	"github.com/ppp3ppj/wywy/pub/pages/signup"
	"github.com/ppp3ppj/wywy/template"
)

func (ue *UserFrontend) GetSignup(c echo.Context) error {
    signup := signup.Signup()

    return template.AssertRender(c, http.StatusOK, signup)
}

func (ue *UserFrontend) Signup(c echo.Context) error {

    registerView := signup.Signup()

    if c.Request().Method == "POST" {
        fmt.Println("POST")
        /*
        user := services.User{
            Email:    c.FormValue("email"),
            Password: c.FormValue("password"),
            Username: c.FormValue("username"),
        }
        */
        user := models.User {
            Email: c.FormValue("email"),
            Password: c.FormValue("password"),
            Username: c.FormValue("username"),
            Firstname: c.FormValue("firstname"),
            Lastname: c.FormValue("lastname"),
        }

        err := ue.userRepo.CreateUser(user)
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
        return c.Redirect(http.StatusSeeOther, "/")
    }

    return template.AssertRender(c, http.StatusOK, registerView)
}


