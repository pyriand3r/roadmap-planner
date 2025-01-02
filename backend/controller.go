package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ###### Feature API ###### //

func featureGetList(ctx echo.Context) error {
	var features []Feature
	result := db.Find(&features)
	if result.Error != nil {
		return result.Error
	}
	return ctx.JSONPretty(http.StatusOK, features, "")
}

func featurePost(ctx echo.Context) error {
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

func featurePut(ctx echo.Context) error {
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

func featureDelete(ctx echo.Context) error {
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

func milestoneGetList(ctx echo.Context) error {
	var milestones []Milestone
	result := db.Find(&milestones)
	if result.Error != nil {
		return result.Error
	}
	return ctx.JSONPretty(http.StatusOK, milestones, "")
}

func milestonePost(ctx echo.Context) error {
	m := new(Milestone)
	if err := ctx.Bind(m); err != nil {
		return err
	}
	result := db.Create(m)
	if result.Error != nil {
		return result.Error
	}
	return ctx.JSONPretty(http.StatusOK, m.ID, "")
}
