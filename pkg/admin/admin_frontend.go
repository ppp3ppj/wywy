package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	pub_index "github.com/ppp3ppj/wywy/pub/pages/index"
	pub_variables "github.com/ppp3ppj/wywy/pub/variables"
	"github.com/ppp3ppj/wywy/template"
)

type AdminFrontend struct {
}

func NewAdminFrontend(
    g *echo.Group,
) {
    fe := &AdminFrontend{}
    g.GET("", fe.Index)
}

func (fe *AdminFrontend) Index(c echo.Context) error {
    bodyOpts := pub_variables.BodyOpts {
        HeaderButtons: nil,
        Component: nil,
    }

    index := pub_index.Index(bodyOpts)

    return template.AssertRender(c, http.StatusOK, index)
}
