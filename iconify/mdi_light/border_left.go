package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4v8h1v1H4v8H3V4h1m2 17v-1h2v1H6m4-16V4h3v1h-1v1h-1V5h-1M6 5V4h2v1H6m5 8v-1h1v1h-1m8-9h1v1h-1V4m-2 0v1h-2V4h2m2 12h1v2h-1v-2m0-9h1v2h-1V7m0 4h1v3h-1v-1h-1v-1h1v-1m-9 9h1v-1h1v1h1v1h-3v-1m5 1v-1h2v1h-2m4-1h1v1h-1v-1m-8-5h1v2h-1v-2m-4-2v-1h2v1H7m7 0v-1h2v1h-2m-3-5h1v2h-1V8Z"/>`),
		g.Group(children),
	)
}