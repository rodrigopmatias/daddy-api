package terminal

import (
	"net/http"

	"github.com/matias-inc/muxapi/router"
	"github.com/rodrigopmatias/daddy-api/db/controllers"
)

func retriveResource(c router.RouterContext) router.RouterResponse {
	terminal, err := controllers.TerminalController.Get(c.Params["id"])
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	return router.DataResponse(http.StatusOK, terminal)
}
