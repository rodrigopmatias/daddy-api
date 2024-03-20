package terminal

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/matias-inc/muxapi/router"
	"github.com/rodrigopmatias/daddy-api/db/controllers"
	"github.com/rodrigopmatias/daddy-api/db/input"
)

func updateResource(c router.RouterContext) router.RouterResponse {
	body := input.Terminal{}
	json.NewDecoder(c.Request.Body).Decode(&body)

	if err := body.IsValid(); err != nil {
		return router.DetailResponse(http.StatusBadRequest, err.Error())
	}

	exists, err := controllers.TerminalController.Exists(controllers.NewFilter("id = ?", c.Params["id"]))
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	if exists {
		if err := controllers.TerminalController.Update(c.Params["id"], body); err != nil {
			return router.DetailResponse(err.StatusCode(), err.Error())
		}

		terminal, err := controllers.TerminalController.Get(c.Params["id"])
		if err != nil {
			return router.DetailResponse(err.StatusCode(), err.Error())
		}

		return router.DataResponse(http.StatusOK, terminal)
	}

	terminal, err := controllers.TerminalController.CreateWithId(c.Params["id"], body)
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	return router.DataResponse(http.StatusCreated, terminal)
}
