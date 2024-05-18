package dashboard

import (
	"net/http"

	"github.com/labstack/echo/v4"
	pub_dashboard_todos "github.com/ppp3ppj/wywy/pub/pages/dashboards/Todos"
	pub_variables "github.com/ppp3ppj/wywy/pub/variables"
	"github.com/ppp3ppj/wywy/template"
)

func (fe *DashboardFrontend) Todos(c echo.Context) error {
    opts := pub_variables.DashboardOpts{ }
    todoVM := pub_dashboard_todos.TodosVM{
        Opts: opts,
    }
    dashboard := pub_dashboard_todos.Todos(todoVM)

    return template.AssertRender(c, http.StatusOK, dashboard)
}
