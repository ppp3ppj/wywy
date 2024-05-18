package dashboard

import "github.com/labstack/echo/v4"


type DashboardFrontend struct {
}

func NewDashboardFrontend(
    g *echo.Group,
    authMid echo.MiddlewareFunc,
) {
    fe := &DashboardFrontend{}

    g.GET("/dashboard", fe.Todos, authMid)
    g.GET("/dashboard/login", fe.Login)
}
