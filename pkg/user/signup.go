package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/pub/pages/signup"
	"github.com/ppp3ppj/wywy/template"
)

func (ue *UserFrontend) Signup(c echo.Context) error {
    signup := signup.Signup()

    return template.AssertRender(c, http.StatusOK, signup)
}
