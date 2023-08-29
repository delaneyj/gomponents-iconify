package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterBoilerAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2c-1.11 0-2 .89-2 2v12c0 1.11.89 2 2 2h1v2H4v2h3c1.11 0 2-.89 2-2v-2h2v2c0 1.11.89 2 2 2h3v-2h-3v-2h1c1.11 0 2-.89 2-2V4c0-1.11-.89-2-2-2H6m4 2.97a2 2 0 1 1 .001 3.999A2 2 0 0 1 10 4.97M8 14.5h4V16H8v-1.5m10 .5h2v2h-2v-2m0-8h2v6h-2V7Z"/>`),
		g.Group(children),
	)
}