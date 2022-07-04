package api

import (
	"net/http"

	"github.com/frankmeza/blog-go-htmx/pkg/blog"
	"github.com/frankmeza/blog-go-htmx/pkg/utils"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

const FOUR_SPACES = "    "

func getAppStateJson(c echo.Context) error {
	appState := blog.BackdoorBootstrap()
	return c.JSONPretty(http.StatusOK, appState, FOUR_SPACES)
}

func getAppStateYaml(c echo.Context) error {
	appState := blog.BackdoorBootstrap()

	yamlAsBytes, err := yaml.Marshal(&appState)
	if err != nil {
		return utils.ReturnError("getAppStateYaml:24", err)
	}

	return c.String(http.StatusOK, string(yamlAsBytes))
}

func AddApiActions(echoServer *echo.Echo) {
	echoServer.GET("/f/app-state.json", getAppStateJson)
	echoServer.GET("/f/app-state.yaml", getAppStateYaml)
}
