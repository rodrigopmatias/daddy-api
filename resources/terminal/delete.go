package terminal

import (
	"github.com/matias-inc/muxapi/router"
	"github.com/rodrigopmatias/daddy-api/db/controllers"
)

func deleteResource(c router.RouterContext) router.RouterResponse {
	if err := controllers.TerminalController.Delete(c.Params["id"]); err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	return router.NoContentResponse()
}
