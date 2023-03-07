package controller

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

func Upload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["image"]
	filename := ""

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		filename = file.Filename

		// Destination
		dst, err := os.Create("./cmd/uploads/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return c.JSON(http.StatusAccepted, echo.Map{"url": "http://localhost:1323/api/v1/s/images/uploads/" + filename})
}
