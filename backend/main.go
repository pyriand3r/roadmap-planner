package main

import (
	"github.com/labstack/echo/v4/middleware"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var workDir string

func main() {
	var err error
	workDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("could not get current working directory:" + err.Error())
	}
	db, err = gorm.Open(sqlite.Open("features.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Feature{})
	db.AutoMigrate(&Milestone{})

	e := echo.New()

	e.Use(
		middleware.StaticWithConfig(
			middleware.StaticConfig{
				Skipper: func(ctx echo.Context) bool {
					if strings.HasPrefix(ctx.Request().URL.String(), "/api/") {
						return true
					}
					return false
				},
				Root:  workDir + "/dist",
				Index: "index.html",
				HTML5: true,
			},
		),
	)

	e.GET("/api/feature", handleGetList)
	e.POST("/api/feature", handlePostFeature)
	e.PUT("/api/feature", handlePutFeature)
	e.DELETE("/api/feature/:id", handleDeleteFeature)

	e.Logger.Fatal(e.Start(":1323"))
}
