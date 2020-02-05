package routes

import (
	"github.com/labstack/echo"
	"goNuxtTube/web/api"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
	}
}