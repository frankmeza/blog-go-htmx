package blog

import (
	"html/template"

	"github.com/frankmeza/blog-go-htmx/pkg/utils"
	"github.com/labstack/echo/v4"
)

var appState AppState

var t = &GoTemplateSet{
	templates: template.Must(template.ParseGlob(
		utils.UseRelativePath("pkg/templates/*.html"),
	)),
}

func AddBlogActions(echoServer *echo.Echo) {
	echoServer.GET("/", bootstrapAppState)
	echoServer.GET("/pages/:page", renderMarkdownAsHtml)
	echoServer.GET("/posts/:post", renderMarkdownAsHtml)
}
