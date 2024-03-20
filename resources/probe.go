package resources

import (
	"net/http"

	"github.com/matias-inc/muxapi/router"
)

func liveResource(c router.RouterContext) router.RouterResponse {
	return router.DetailResponse(http.StatusOK, "Im live!!!")
}

func readyResource(c router.RouterContext) router.RouterResponse {
	return router.DetailResponse(http.StatusOK, "Im ready!!!")
}
