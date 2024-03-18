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
    sess, _ := session.Get(auth_sessios_key, c)

    uNav := models.UserNav{
        Id: sess.Values["user_id"].(string),
        Username: sess.Values["username"].(string),
    }

    fmt.Println(sess.Values["user_id"])
    return renderView(c, dashboard_views.DashboardIndex(
        uNav,
        fromProtected,
        dashboard_views.DashboardList(),
    ))
}
