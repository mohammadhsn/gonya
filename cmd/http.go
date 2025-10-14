package cmd

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		startHttp(cmd.Context())
	},
}

func startHttp(_ context.Context) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
