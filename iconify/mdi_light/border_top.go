package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 5h-8v1h-1V5H3V4h17v1M3 7h1v2H3V7m16 4h1v3h-1v-1h-1v-1h1v-1m0-4h1v2h-1V7m-8 5h1v1h-1v-1m9 8v1h-1v-1h1m0-2h-1v-2h1v2M8 20v1H6v-1h2m9 0v1h-2v-1h2m-4 0v1h-3v-1h1v-1h1v1h1m-9-9v1h1v1H4v1H3v-3h1m-1 5h1v2H3v-2m1 4v1H3v-1h1m5-8v1H7v-1h2m2-4h1v2h-1V8m0 7h1v2h-1v-2m5-3v1h-2v-1h2Z"/>`),
		g.Group(children),
	)
}