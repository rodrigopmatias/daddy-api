package terminal

import (
	"net/http"

	"github.com/matias-inc/muxapi/router"
	"github.com/rodrigopmatias/daddy-api/db/controllers"
)

func listResource(c router.RouterContext) router.RouterResponse {
	count, err := controllers.TerminalController.Count()
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	items, err := controllers.TerminalController.List(0, 20)
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	return router.DataResponse(http.StatusOK, router.H{
		"next":    nil,
		"previus": nil,
		"count":   count,
		"items":   items,
	})
}
