package wywyauth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/pkg/models"
)

const (
    auth_sessios_key string = "authenticate-sessions"
    auth_key string = "authenticated"
    user_id_key string = "user_id"
    username_key string = "username"
    tz_key string = "time_zone"
)

func WywyAuthMiddleware(userRepo models.UserRepository) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            sess, err := session.Get("test_token", c)
            if auth, ok := sess.Values[auth_key].(bool); !ok || !auth {
                fmt.Println("agent--")
                fmt.Println(auth)
                //return echo.NewHTTPError(echo.ErrUnauthorized.Code, "Please provide valid credentials")
                return c.Redirect(http.StatusFound, "/dashboard/login")
            }

            if userId, ok := sess.Values[user_id_key].(string); !ok || len(userId) == 0 {
                c.Set(user_id_key, userId)
            }

            if username, ok := sess.Values[username_key].(string); !ok || len(username) == 0 {
                c.Set(username_key, username)
            }
            if err != nil {
                return c.Redirect(http.StatusFound, "/")
            }
            fmt.Println("from middleware------")
            fmt.Println(sess.Values[auth_key].(bool))

            return next(c)
        }
    }
}

