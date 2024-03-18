package handlers

import (
	"fmt"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/models"
	"github.com/ppp3ppj/wywy/views/dashboard_views"
)

type DashboardService interface {
}

type DashboardHandler struct {
    DashboardService DashboardService
}

func NewDashboardHandler(ds DashboardService) *DashboardHandler {
    return &DashboardHandler{
        DashboardService: ds,
    }
}

func (h *DashboardHandler) dashboardListHandler(c echo.Context) error {
    // decrypt the user id from the session
    // get the user's dashboard dashboardListHandler

    sess, _ := session.Get(auth_sessios_key, c)
    fmt.Println(sess.Values["user_id"])
    userNav := models.UserNav{}
    return renderView(c, dashboard_views.DashboardIndex(
        userNav,
        fromProtected,
        dashboard_views.DashboardList(),
    ))
}
