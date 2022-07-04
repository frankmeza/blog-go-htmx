package assets

import (
	"github.com/frankmeza/blog-go-htmx/pkg/utils"
	"github.com/labstack/echo/v4"
)

func AddStaticAssets(echoServer *echo.Echo) {
	echoServer.Static("/assets/", utils.UseRelativePath("pkg/assets"))
}
