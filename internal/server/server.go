package server

import (
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/config"
	"github.com/ppp3ppj/wywy/db"
)


type echoServer struct {
    app *echo.Echo
    conf *config.Config
    db db.IDatabase
}
