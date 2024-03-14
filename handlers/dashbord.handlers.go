package handlers

import (
	"github.com/labstack/echo/v4"
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
    return renderView(c, dashboard_views.DashboardIndex(
        "",
        fromProtected,
        dashboard_views.DashboardList(),
    ))
}
