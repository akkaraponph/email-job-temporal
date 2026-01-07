package routers

import "gofr.dev/pkg/gofr"

type RouterImpls struct {
	app *gofr.App
}

func NewRoute(app *gofr.App) RouterImpls {
	return RouterImpls{app: app}
}
