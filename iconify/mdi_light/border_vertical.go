package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4v1h-2V4h2m-4 16v1h-3v-1h1V5h-1V4h3v1h-1v15h1m4 0v1h-2v-1h2M4 21H3v-1h1v1m2 0v-1h2v1H6M4 9H3V7h1v2m0 9H3v-2h1v2m0-4H3v-3h1v1h1v1H4v1M8 4v1H6V4h2M4 5H3V4h1v1m12 7v1h-2v-1h2m-7 0v1H7v-1h2m10 9v-1h1v1h-1m0-12V7h1v2h-1m0 9v-2h1v2h-1m0-4v-1h-1v-1h1v-1h1v3h-1m0-9V4h1v1h-1Z"/>`),
		g.Group(children),
	)
}