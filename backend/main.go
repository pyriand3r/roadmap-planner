package main

import (
	"net/http"

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
	db.AutoMigrate(&Milestone{})

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
