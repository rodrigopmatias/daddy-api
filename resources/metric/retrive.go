package metric

import (
	"net/http"

	"github.com/matias-inc/muxapi/router"
	"github.com/rodrigopmatias/daddy-api/db/controllers"
)

func retriveResource(c router.RouterContext) router.RouterResponse {
	metric, err := controllers.MetricController.Get(c.Params["id"])
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	return router.DataResponse(http.StatusOK, metric)
}
