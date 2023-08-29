package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartHistogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h1v9h3V7h5v4h4v4h4v6H3V4m13 12v4h3v-4h-3m-4-4v8h3v-8h-3M8 8v12h3V8H8m-4 6v6h3v-6H4Z"/>`),
		g.Group(children),
	)
}