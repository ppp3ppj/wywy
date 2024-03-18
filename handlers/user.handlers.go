package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/wywy/models"
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
    userProfile := models.UserProfile{
        Id: "1",
        Email: "uuu@mail.com",
        Username: "uuu",
        Firstname: "uuu",
        Lastname: "uuu",
    }
    profileView := user_profile_views.UserProfileSetting(userProfile)

    userNav := models.UserNav{}

    return renderView(c, user_profile_views.UserProfileSettingIndex(
        userNav,
        fromProtected,
        profileView,
    ))
}
