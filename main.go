package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/codescot/go-common/fileio"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	filepath.Walk("api", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".json") {
			route := path[4 : len(path)-5]

			r.GET(route, func(ctx *gin.Context) {
				file := make(map[string]interface{})
				fileio.ReadJSON(path, &file)

				ctx.JSON(http.StatusOK, file)
			})
		}

		return nil
	})

	r.Run(":9000")
}
