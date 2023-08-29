package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taxi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M112 304v-96h40v-32H40v32h40v96h32zm221.483 0L356 269.358L378.517 304h38.166l-41.6-64l41.6-64h-38.166L356 210.642L333.483 176h-38.166l41.6 64l-41.6 64h38.166zM440 176h32v128h-32zM40 104h432v32H40zm0 240h432v32H40zm201.337-64l8 24h33.731L240.4 176h-45.952l-39.439 128h33.484l7.4-24Zm-23.617-70.854L230.671 248h-24.923Z"/>`),
		g.Group(children),
	)
}