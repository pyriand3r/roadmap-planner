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

	if _, err := os.Stat(workDir + "/data"); os.IsNotExist(err) {
		err = os.Mkdir(workDir+"/data", 0755)
		if err != nil {
			panic("could not create data directory:" + err.Error())
		}
	}

	db, err = gorm.Open(sqlite.Open("data/features.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database:")
	}
	if err := db.AutoMigrate(&Feature{}); err != nil {
		panic("failed to migrate features:" + err.Error())
	}
	if err := db.AutoMigrate(&Milestone{}); err != nil {
		panic("failed to migrate milestones:" + err.Error())
	}

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
