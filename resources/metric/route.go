package metric

import "github.com/matias-inc/muxapi/router"

func InitRouter(r *router.RouterGroup) {
	blueprint := router.NewRouterGroup("/metrics")

	blueprint.Post("", createResource)
	blueprint.Get("", listResource)

	blueprint.Get("/{id:uuid}", retriveResource)

	r.IncludeRouterGroup(blueprint)
}
