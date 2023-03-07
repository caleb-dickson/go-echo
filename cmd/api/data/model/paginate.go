package model

import (
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/data/model/entity"
)

func Paginate(c echo.Context, ent entity.Entity, page int) echo.Map {

	limit := 15
	offset := (page - 1) * limit

	entities := ent.Take(DB, limit, offset)
	totalEntities := ent.Count(DB)

	return echo.Map{
		"meta": echo.Map{
			"totalResults": totalEntities,
			"currentPage":  page,
			"lastPage":     float64(totalEntities/int64(limit)) + 1,
		},
		"data": entities,
	}
}
