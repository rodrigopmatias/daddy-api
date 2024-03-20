package terminal

import "github.com/matias-inc/muxapi/router"

func InitRouter(r *router.RouterGroup) {
	blueprint := router.NewRouterGroup("/terminals")

	blueprint.Post("", createResource)
	blueprint.Get("", listResource)

	blueprint.Get("/{id:uuid}", retriveResource)
	blueprint.Patch("/{id:uuid}", partialUpdateResource)
	blueprint.Put("/{id:uuid}", updateResource)
	blueprint.Delete("/{id:uuid}", deleteResource)

	r.IncludeRouterGroup(blueprint)
}
