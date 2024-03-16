package handlers

import (
	"github.com/labstack/echo/v4"
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
    profileView := user_profile_views.UserProfileSetting()
    return renderView(c, user_profile_views.UserProfileSettingIndex(
        "",
        fromProtected,
        profileView,
    ))
}
