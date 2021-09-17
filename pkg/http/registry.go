package http

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ApiV2 = "/api/v2"
const ApiMeta = "/-"

func registerHandlers(router *gin.Engine, capturesRepo captures.Repository) {
	assignHandlers(router.Group(ApiMeta), initApiMetaRoutes(capturesRepo))
	assignHandlers(router.Group(ApiV2), initApiV2Routes(capturesRepo))
}

func assignHandlers(routerGroup *gin.RouterGroup, routes []Route) {
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			routerGroup.GET(route.Pattern, route.Handler.Handle)
		case http.MethodPost:
			routerGroup.POST(route.Pattern, route.Handler.Handle)
		case http.MethodPut:
			routerGroup.PUT(route.Pattern, route.Handler.Handle)
		case http.MethodPatch:
			routerGroup.PATCH(route.Pattern, route.Handler.Handle)
		case http.MethodDelete:
			routerGroup.DELETE(route.Pattern, route.Handler.Handle)
		}
	}
}

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string

	Handler Handler
}

// Routes is the list of the generated Route.
type Routes []Route
