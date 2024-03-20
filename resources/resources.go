package resources

import "github.com/matias-inc/muxapi/router"

func InitRouter(base *router.RouterGroup) {
	base.Get("/live", liveResource)
	base.Get("/ready", readyResource)
}
