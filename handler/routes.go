package handler

import (
	"github.com/arvnd79/ie-proj/common"
	"github.com/arvnd79/ie-proj/middleware"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
)

// RegisterRoutes registers routes with their corresponding handler function
// functions are defined in handler package
func (h *Handler) RegisterRoutes(v *echo.Group) {

	v.Use(middleware.JWT(common.JWTSecret))
	v.Use(middleware2.RemoveTrailingSlash())

	// adding white list
	middleware.AddToWhiteList("/api/users/login", "POST")
	middleware.AddToWhiteList("/api/users", "POST")

	userGroup := v.Group("/users")
	userGroup.POST("", h.SignUp)
	userGroup.POST("/login", h.Login)

	urlGroup := v.Group("/urls")
	urlGroup.GET("", h.FetchURLs)
	urlGroup.POST("", h.CreateURL)
	urlGroup.GET("/:urlID", h.GetURLStats)

	alertGroup := v.Group("/alerts")
	alertGroup.GET("", h.FetchAlerts)

}
