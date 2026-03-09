package server

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	_ "isOdin/RestApi/pkg/api/swagger"
)

type HandlerInterface interface {
	// Auth
	SignInHandler(c *echo.Context) error
	SignUpHandler(c *echo.Context) error

	// Item
	CreateItem(c *echo.Context) error
	GetItem(c *echo.Context) error
	UpdateItem(c *echo.Context) error
	DeleteItem(c *echo.Context) error

	// List
	CreateList(c *echo.Context) error
	GetList(c *echo.Context) error
	UpdateList(c *echo.Context) error
	DeleteList(c *echo.Context) error
}

type MiddlewareInterface interface {
	JWTAuth() echo.MiddlewareFunc
}

func NewRouter(e *echo.Echo, md MiddlewareInterface, h HandlerInterface) {
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	api := e.Group("/api/v0")

	auth := api.Group("/auth")
	auth.POST("/sign-in", h.SignInHandler)
	auth.POST("/sign-up", h.SignUpHandler)

	list := api.Group("/list")
	list.Use(md.JWTAuth())

	list.POST("/", h.CreateList)
	list.GET("/:list_id", h.GetList)
	list.PATCH("/:list_id", h.UpdateList)
	list.DELETE("/:list_id", h.DeleteList)

	item := list.Group("/:list_id/item")
	item.POST("/", h.CreateItem)
	item.GET("/:item_id", h.GetItem)
	item.PATCH("/:item_id", h.UpdateItem)
	item.DELETE("/:item_id", h.DeleteItem)
}
