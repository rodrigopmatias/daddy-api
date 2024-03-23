package terminal

import (
	"net/http"

	"github.com/matias-inc/muxapi/router"
	"github.com/rodrigopmatias/daddy-api/db/controllers"
	"github.com/rodrigopmatias/daddy-api/helpers"
)

func listResource(c router.RouterContext) router.RouterResponse {
	count, err := controllers.TerminalController.Count()
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	offset, limit := helpers.ExtractPaginateInfo(c.Request.URL)
	order := make([]string, 0)
	for _, field := range helpers.ExtractOrdering(c.Request.URL) {
		order = append(order, field.String())
	}

	items, err := controllers.TerminalController.List(offset, limit, order)
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	nextURL, previusURL := helpers.Paginate(c.Request.URL, count, offset, limit)

	return router.DataResponse(http.StatusOK, router.H{
		"next":    nextURL,
		"previus": previusURL,
		"count":   count,
		"items":   items,
	})
}
