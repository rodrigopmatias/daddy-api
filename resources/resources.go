package resources

import (
	"github.com/matias-inc/muxapi/router"
	"github.com/rodrigopmatias/daddy-api/resources/terminal"
)

func InitRouter(base *router.RouterGroup) {
	base.Get("/live", liveResource)
	base.Get("/ready", readyResource)

	terminal.InitRouter(base)
}
