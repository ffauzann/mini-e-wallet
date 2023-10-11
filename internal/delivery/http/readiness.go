package http

import "github.com/labstack/echo/v4"

// @Summary	Health check endpoint for k8s
// @Tags	Health and Readiness
// @Accept	json
// @Produce	text/plain
// @Success	200
// @Router	/health [get]
func health(c echo.Context) error {
	return c.String(200, "healthy upstream")
}

// @Summary	Readiness endpoint for k8s
// @Tags	Health and Readiness
// @Accept	json
// @Produce	text/plain
// @Success	200
// @Router	/readiness [get]
func readiness(c echo.Context) error {
	return c.String(200, "OK")
}
