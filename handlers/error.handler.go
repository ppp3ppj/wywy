package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/views/errors_pages"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
    code := http.StatusInternalServerError
    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
    }

    c.Logger().Error(err)

    var errorPage func(fp bool) templ.Component

    switch code { 
    case 500: // Internal Server errorPage
        errorPage = errors_pages.Error500
    }
    isError = true

    renderView(c, errors_pages.ErorrIndex(
        isError,
        errorPage(false),
    ))
}
