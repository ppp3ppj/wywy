package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ppp3ppj/wywy/config"
	"github.com/ppp3ppj/wywy/db"
	"github.com/ppp3ppj/wywy/services"

	"github.com/labstack/echo/v4"
//	"github.com/labstack/echo/v4/middleware"
)

func envPath() string {
    if len(os.Args) == 1 {
        return ".env"
    } else {
        return os.Args[1]
    }
}

func main() {
    e := echo.New()
    e.Static("/", "assets")
    e.Use(loggerMiddleware)
    cfg := config.LoadConfig(envPath())
    db := db.DbConnect(cfg.Db())

    us := services.NewUserService(services.User{}, db)
    _ = us
    e.Logger.Fatal(e.Start(":1323"))
}

func loggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Log incoming request
        log.Printf("Incoming Request: %s %s", c.Request().Method, c.Request().URL.String())
        // Call next handler
        if err := next(c); err != nil {
            c.Error(err)
        }
        // Log outgoing response
        log.Printf("Outgoing Response: %d %s", c.Response().Status, http.StatusText(c.Response().Status))
        return nil
    }
}
