package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2.5L20 7v2H2V7l9-4.5m4 7.5h4v8h-4v-8M2 22v-3h18v3H2m7-12h4v8H9v-8m-6 0h4v8H3v-8m0 10v1h16v-1H3m1-9v6h2v-6H4m6 0v6h2v-6h-2m6 0v6h2v-6h-2M3 8h16v-.4l-8-4.02L3 7.6V8Z"/>`),
		g.Group(children),
	)
}