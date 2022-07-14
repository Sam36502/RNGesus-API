package api

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/Sam36502/RNGesus-API/divinity"
	"github.com/Sam36502/RNGesus-API/dto"
	"github.com/labstack/echo/v4"
)

const (
	PARAM_INT_BASE        = 10
	PARAM_INT_BITS        = 64
	PARAM_INT_MIN_DEFAULT = math.MinInt64
	PARAM_INT_MAX_DEFAULT = math.MaxInt64
	FAILURE_STATUS        = 666
)

func getRandomFloat(c echo.Context) error {
	num, err := divinity.RandomFloat()
	if err != nil {
		return c.JSON(
			FAILURE_STATUS,
			dto.ErrorResponse{
				Message: err.Error(),
			},
		)
	}
	return c.JSON(http.StatusOK, dto.FloatResponse{
		Number: num,
	})
}

func getRandomInt(c echo.Context) error {

	// Get range parameters
	minStr := c.QueryParam(dto.PARAM_INT_MIN)
	maxStr := c.QueryParam(dto.PARAM_INT_MAX)

	// Parse minimum parameter
	var err error
	var min int64
	if minStr == "" {
		min = PARAM_INT_MIN_DEFAULT
	} else {
		min, err = strconv.ParseInt(minStr, PARAM_INT_BASE, PARAM_INT_BITS)
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				dto.ErrorResponse{
					Message: fmt.Sprintf("Error: Invalid `min` parameter ('%s') provided:\n%s", minStr, err.Error()),
				},
			)
		}
	}

	// Parse maximum parameter
	var max int64
	if maxStr == "" {
		max = PARAM_INT_MAX_DEFAULT
	} else {
		max, err = strconv.ParseInt(maxStr, PARAM_INT_BASE, PARAM_INT_BITS)
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				dto.ErrorResponse{
					Message: fmt.Sprintf("Error: Invalid `max` parameter ('%s') provided:\n%s", maxStr, err.Error()),
				},
			)
		}
	}

	num, err := divinity.RandomInt(min, max)
	if err != nil {
		return c.JSON(
			FAILURE_STATUS,
			dto.ErrorResponse{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, dto.IntResponse{
		Number: num,
	})
}
