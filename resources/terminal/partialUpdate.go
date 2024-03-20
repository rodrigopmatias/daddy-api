package terminal

import (
	"net/http"

	"github.com/matias-inc/muxapi/router"
)

func partialUpdateResource(c router.RouterContext) router.RouterResponse {
	return router.DetailResponse(http.StatusNotImplemented, "not implemented")
}
