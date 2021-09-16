package http

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"github.com/gin-gonic/gin"
	"net/http"
)

const ApiV2 = "/api/v2"

func registerHandlers(router *gin.Engine, capturesRepo captures.Repository) {
	v2 := router.Group(ApiV2)

	routes := initRoutes(capturesRepo)

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			v2.GET(route.Pattern, route.Handler.Handle)
		case http.MethodPost:
			v2.POST(route.Pattern, route.Handler.Handle)
		case http.MethodPut:
			v2.PUT(route.Pattern, route.Handler.Handle)
		case http.MethodPatch:
			v2.PATCH(route.Pattern, route.Handler.Handle)
		case http.MethodDelete:
			v2.DELETE(route.Pattern, route.Handler.Handle)
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
