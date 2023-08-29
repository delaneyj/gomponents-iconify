package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOutside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h17v17H3V4m1 1v7h1v1H4v7h7v-1h1v1h7v-7h-1v-1h1V5h-7v1h-1V5H4m3 7h2v1H7v-1m4 0h1v1h-1v-1m0 3h1v2h-1v-2m3-3h2v1h-2v-1m-3-2V8h1v2h-1Z"/>`),
		g.Group(children),
	)
}