package api

import (
	"net/http"

	"github.com/ymsht/mosaic-backend/model"
	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetArts 芸術情報を返します
func GetArts() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		arts, err := model.GetArts(tx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, arts)
	}
}
