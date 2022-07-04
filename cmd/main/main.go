package main

import (
	"flag"

	"github.com/frankmeza/blog-go-htmx/pkg/api"
	"github.com/frankmeza/blog-go-htmx/pkg/assets"
	blog "github.com/frankmeza/blog-go-htmx/pkg/blog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	isDev := flag.Int(
		"dev", 0, "append -dev 1 to start server on :9990",
	)

	flag.Parse()
	server := echo.New()

	if *isDev != 0 {
		server.Use(middleware.Logger())
		server.Use(middleware.Recover())
	}

	api.AddApiActions(server)
	assets.AddStaticAssets(server)
	blog.AddBlogActions(server)

	server.Logger.Fatal(server.Start("127.0.0.1:9990"))
}
