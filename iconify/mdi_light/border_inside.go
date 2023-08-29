package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderInside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4v1h-2V4h2m-4 16v1h-3v-1h1v-7H4v1H3v-3h1v1h7V5h-1V4h3v1h-1v7h7v-1h1v3h-1v-1h-7v7h1m4 0v1h-2v-1h2M4 21H3v-1h1v1m2 0v-1h2v1H6M4 9H3V7h1v2m0 9H3v-2h1v2M8 4v1H6V4h2M4 5H3V4h1v1m15 16v-1h1v1h-1m0-12V7h1v2h-1m0 9v-2h1v2h-1m0-13V4h1v1h-1Z"/>`),
		g.Group(children),
	)
}