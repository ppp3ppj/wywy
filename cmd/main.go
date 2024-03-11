package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/ppp3ppj/wywy/config"
	"github.com/ppp3ppj/wywy/db"
	"github.com/ppp3ppj/wywy/handlers"
	"github.com/ppp3ppj/wywy/services"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//	"github.com/labstack/echo/v4/middleware"
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
    //e.Static("/", "assets")
    e.Static("/", "public")
    e.Use(loggerMiddleware)

    // Cors Middleware
    e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

    // e.Use(middleware.Logger())
    cfg := config.LoadConfig(envPath())

    // Session middleware
    // Secret Key is JWT_ADMIN_KEY at env.dev
    fmt.Println("Secret Key: ", cfg.Jwt().SecretKey())
    e.Use(session.Middleware(sessions.NewCookieStore([]byte(cfg.Jwt().SecretKey()))))
    db := db.DbConnect(cfg.Db())

    us := services.NewUserService(services.User{}, db)
    _ = us

    ah := handlers.NewAuthHandler(us)
    handlers.SetupRoutes(e, ah)


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

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
