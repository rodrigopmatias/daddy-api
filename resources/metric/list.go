package metric

import (
	"net/http"

	"github.com/matias-inc/muxapi/router"
	"github.com/rodrigopmatias/daddy-api/db/controllers"
	"github.com/rodrigopmatias/daddy-api/helpers"
)

var filterMap = controllers.ExtractFilterMap{
	"before": func(key, value string) controllers.Filter {
		return controllers.NewFilter("created_at <= ?", value)
	},
	"after": func(key, value string) controllers.Filter {
		return controllers.NewFilter("created_at >= ?", value)
	},
}

func listResource(c router.RouterContext) router.RouterResponse {
	filters := controllers.ExtractFilters(c.Request.URL, filterMap)

	count, err := controllers.MetricController.Count(filters...)
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	offset, limit := helpers.ExtractPaginateInfo(c.Request.URL)
	order := make([]string, 0)
	for _, field := range helpers.ExtractOrdering(c.Request.URL) {
		order = append(order, field.String())
	}

	items, err := controllers.MetricController.List(offset, limit, order, filters)
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	nextURL, previusURL := helpers.Paginate(c.Request.URL, count, offset, limit)

	return router.DataResponse(http.StatusOK, router.H{
		"next":    previusURL,
		"previus": nextURL,
		"count":   count,
		"items":   items,
	})
}
