package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 21v-8h-1v-1h1V4h1v17h-1M17 4v1h-2V4h2m-4 16v1h-3v-1h1v-1h1v1h1m4 0v1h-2v-1h2m-5-8v1h-1v-1h1m-8 9H3v-1h1v1m2 0v-1h2v1H6M4 9H3V7h1v2m0 9H3v-2h1v2m0-4H3v-3h1v1h1v1H4v1m9-9h-1v1h-1V5h-1V4h3v1M8 4v1H6V4h2M4 5H3V4h1v1m8 5h-1V8h1v2m4 2v1h-2v-1h2m-7 0v1H7v-1h2m3 5h-1v-2h1v2Z"/>`),
		g.Group(children),
	)
}