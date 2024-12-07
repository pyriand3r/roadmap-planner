package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("features.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Feature{})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/feature", handleGetList)
	e.POST("/api/feature", handlePostFeature)
	e.PUT("/api/feature", handlePutFeature)
	e.DELETE("/api/feature/:id", handleDeleteFeature)

	e.Logger.Fatal(e.Start(":1323"))
}

func handleGetList(ctx echo.Context) error {
	features := []Feature{}
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
