package dashboard

import (
	"net/http"

	"github.com/labstack/echo/v4"
	pub_dashboard_login "github.com/ppp3ppj/wywy/pub/pages/dashboards/login"
	pub_variables "github.com/ppp3ppj/wywy/pub/variables"
	"github.com/ppp3ppj/wywy/template"
)

func (fe *DashboardFrontend) Login(c echo.Context) error {
    opts := pub_variables.DashboardOpts{
    }

    loginVM := pub_dashboard_login.LoginVM{
        Opts: opts,
    }

    login := pub_dashboard_login.Login(&loginVM)

    return template.AssertRender(c, http.StatusOK, login)
}
