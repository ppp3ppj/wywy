package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/ppp3ppj/wywy/config"
	"github.com/ppp3ppj/wywy/db"
	"github.com/ppp3ppj/wywy/handlers"
	"github.com/ppp3ppj/wywy/pkg/admin"
	"github.com/ppp3ppj/wywy/pkg/auth"
	"github.com/ppp3ppj/wywy/pkg/auth/user_repository"
	"github.com/ppp3ppj/wywy/pkg/dashboard"
	"github.com/ppp3ppj/wywy/pkg/user"
	"github.com/ppp3ppj/wywy/services"
	"github.com/ppp3ppj/wywy/template"
	"github.com/ppp3ppj/wywy/utils/wywyauth"
)


type echoServer struct {
    app *echo.Echo
    conf *config.Config
    db db.IDatabase
}

var (
    server *echoServer
    once sync.Once
)

func NewEchoServer(conf *config.Config, db db.IDatabase) *echoServer {
    echoApp := echo.New()
    echoApp.Logger.SetLevel(log.DEBUG)

    once.Do(func() {
        server = &echoServer{
            app: echoApp,
            conf: config.ConfigGetting(),
            db: db,
        }
    })

    return server
}

func (s *echoServer) Start() {
    timeOutMiddleware := getTimeOutMiddleware(s.conf.Server.Timeout)
    corsMiddleware := getCORSMiddleware(s.conf.Server.AllowOrigins)
    //htmxMiddleware := middlewares.HTMXRequest()

    //s.app.Use(htmxMiddleware)

    s.app.Use(middleware.Recover())
    s.app.Use(middleware.Logger())

    s.app.Use(timeOutMiddleware)
    s.app.Use(corsMiddleware)

    s.app.Static("/dist", "dist")
    s.app.Static("/assets", "public/assets")

    s.app.GET("/v1/health", s.healthCheck)

    // Register template
    template.NewTemplateRenderer(s.app)

    adminGroup := s.app.Group("")
    admin.NewAdminFrontend(adminGroup)


    userRepo := user_repository.NewUserRepository(s.db.Connect())
    authMid := wywyauth.WywyAuthMiddleware(userRepo)
    user.NewUserFrontend(adminGroup, userRepo)
    auth.NewAuthService(adminGroup, userRepo)
    dashboard.NewDashboardFrontend(adminGroup, authMid)

    //s.app.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

    fmt.Println("Secret Key: ", s.conf.AppInfo.SecretKey)
    s.app.Use(session.Middleware(sessions.NewCookieStore([]byte(s.conf.AppInfo.SecretKey))))

    us := services.NewUserService(services.User{}, s.db.Connect())
    ah := handlers.NewAuthHandler(us)

    ds := services.NewDashboardService()
    dh := handlers.NewDashboardHandler(ds)

    uh := handlers.NewUserHandler()

    handlers.SetupRoutes(s.app, ah, dh, uh)

    // Graceful Shutdown
    quitCh := make(chan os.Signal, 1)
    signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
    go s.gracefullyShutdown(quitCh)

    s.httpListening()
}

func (s *echoServer) gracefullyShutdown(quitCh <-chan os.Signal) {
    ctx := context.Background()

    <-quitCh
    s.app.Logger.Info("Shutting down the service...")

    if err := s.app.Shutdown(ctx); err != nil {
        s.app.Logger.Fatalf("Error: %s", err.Error())
    }
}

func (s *echoServer) httpListening() {
    url := fmt.Sprintf(":%d", s.conf.Server.Port)
    if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
        s.app.Logger.Fatalf("shutting down the server: %v", err)
    }
}

func (s *echoServer) healthCheck(c echo.Context) error {
    return c.String(http.StatusOK, "OK")
}

func getTimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
    return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
        Skipper: middleware.DefaultSkipper,
        ErrorMessage: "Error: Request timeout.",
        Timeout: timeout * time.Second,
    })
}

func getCORSMiddleware(allowOrigins []string) echo.MiddlewareFunc {
    allowOrigins = append(allowOrigins, "*://localhost:*")
    return middleware.CORSWithConfig(middleware.CORSConfig{
        Skipper: middleware.DefaultSkipper,
        AllowOrigins: allowOrigins,
        AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    })
}
