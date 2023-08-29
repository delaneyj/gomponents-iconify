package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20h8v-1h1v1h8v1H3v-1m17-2h-1v-2h1v2M4 14H3v-3h1v1h1v1H4v1m0 4H3v-2h1v2m8-5h-1v-1h1v1M3 5V4h1v1H3m0 2h1v2H3V7m12-2V4h2v1h-2M6 5V4h2v1H6m4 0V4h3v1h-1v1h-1V5h-1m9 9v-1h-1v-1h1v-1h1v3h-1m1-5h-1V7h1v2m-1-4V4h1v1h-1m-5 8v-1h2v1h-2m-2 4h-1v-2h1v2m0-7h-1V8h1v2m-5 3v-1h2v1H7Z"/>`),
		g.Group(children),
	)
}