package auth

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/pkg/models"
	pub_dashboard_login "github.com/ppp3ppj/wywy/pub/pages/dashboards/login"
	pub_variables "github.com/ppp3ppj/wywy/pub/variables"
	"github.com/ppp3ppj/wywy/template"
	"golang.org/x/crypto/bcrypt"
)

const (
    auth_sessios_key string = "authenticate-sessions"
    auth_key string = "authenticated"
    user_id_key string = "user_id"
    username_key string = "username"
    tz_key string = "time_zone"
)

type AuthService struct {
    Repo models.UserRepository
}

func NewAuthService(
    g *echo.Group,
    Repo models.UserRepository,
) {
    handler := &AuthService{
        Repo: Repo,
    }

    g.POST("/login", handler.Login)
        // obtaining the time zone from POST Request
}

func (h *AuthService) Login(c echo.Context) error {
    opts := pub_variables.DashboardOpts{
    }

    loginVM := pub_dashboard_login.LoginVM{
        Opts: opts,
    }

    login := pub_dashboard_login.Login(&loginVM)

    if c.Request().Method == "POST" {

        email := c.FormValue("email")
        //password := c.FormValue("password")
        user, err := h.Repo.CheckEmail(email)
        fmt.Println("user is ", user)
        if err != nil {
            if strings.Contains(err.Error(), "email not found") {
                return c.Redirect(http.StatusSeeOther, "/login")
            }
            return echo.NewHTTPError(
                echo.ErrInternalServerError.Code,
                fmt.Sprintf("Failed to check email: %v", err),
            )
        }

        err = bcrypt.CompareHashAndPassword(
            []byte(user.Password),
            []byte(c.FormValue("password")),
        )

        if err != nil {
            return c.Redirect(http.StatusSeeOther, "/login")
        }

        // Get Session and seitting Cookies
        sess, _ := session.Get("test_token", c)
        sess.Options = &sessions.Options{
            Path: "/",
            MaxAge: 3600, // in seconds
            HttpOnly: true,
            Secure: true,
            SameSite: http.SameSiteLaxMode,
        }

        // Set user as authenticated

        sess.Values = map[interface{}]interface{}{
            auth_key: true,
            user_id_key: user.Id,
            username_key: user.Username,
        }
        sess.Save(c.Request(), c.Response())

        return c.Redirect(http.StatusSeeOther, "/")
    //return renderView(c, dashboard_views.DashboardIndex(dashboard_views.DashboardList()))
    }

    return template.AssertRender(c, http.StatusOK, login)
}
