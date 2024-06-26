package handlers

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/models"
	"github.com/ppp3ppj/wywy/services"
	"github.com/ppp3ppj/wywy/views/user_profile_views"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
    return &UserHandler{}
}

type UserService interface {
}

func (h *UserHandler) profileHomeHandler(c echo.Context) error {
    sess, _ := session.Get(auth_sessios_key, c)

    uNav := models.UserNav{
        Id: sess.Values["user_id"].(string),
        Username: sess.Values["username"].(string),
    }


    profileView := user_profile_views.UserProfileSetting(services.User{})
    return renderView(c, user_profile_views.UserProfileSettingIndex(
        uNav,
        fromProtected,
        profileView,
    ))
}
