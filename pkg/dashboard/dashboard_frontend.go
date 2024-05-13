package dashboard

import "github.com/labstack/echo/v4"


type DashboardFrontend struct {
}

func NewDashboardFrontend(
    g *echo.Group,
) {
    fe := &DashboardFrontend{}

    g.GET("/dashboard/login", fe.Login)
}
