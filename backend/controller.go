package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ###### Feature API ###### //

func handleGetList(ctx echo.Context) error {
	var features []Feature
	result := db.Find(&features)
	if result.Error != nil {
		return result.Error
	}
	return ctx.JSONPretty(http.StatusOK, features, "")
}

func handlePostFeature(ctx echo.Context) error {
	f := new(Feature)
	if err := ctx.Bind(f); err != nil {
		return err
	}
	result := db.Create(f)
	if result.Error != nil {
		return result.Error
	}
	return ctx.JSONPretty(http.StatusOK, f.ID, "")
}

func handlePutFeature(ctx echo.Context) error {
	f := new(Feature)
	if err := ctx.Bind(f); err != nil {
		return err
	}
	result := db.Save(f)
	if result.Error != nil {
		return result.Error
	}
	return ctx.JSONPretty(http.StatusOK, f.ID, "")
}

func handleDeleteFeature(ctx echo.Context) error {
	id := ctx.Param("id")
	f := new(Feature)
	uint64ID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return err
	}
	f.ID = uint(uint64ID)
	result := db.Delete(&f)
	if result.Error != nil {
		return result.Error
	}
	return ctx.NoContent(http.StatusNoContent)
}

// ###### Milestone API ###### //
