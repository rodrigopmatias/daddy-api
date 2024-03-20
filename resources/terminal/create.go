package terminal

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/matias-inc/muxapi/router"
	"github.com/rodrigopmatias/daddy-api/db/controllers"
	"github.com/rodrigopmatias/daddy-api/db/input"
)

func createResource(c router.RouterContext) router.RouterResponse {
	body := input.Terminal{}
	json.NewDecoder(c.Request.Body).Decode(&body)

	terminal, err := controllers.TerminalController.Create(body)
	if err != nil {
		return router.DetailResponse(err.StatusCode(), err.Error())
	}

	return router.DataResponse(http.StatusCreated, terminal)
}
