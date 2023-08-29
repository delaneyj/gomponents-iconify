package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Artboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 9v6H7V9h10m2-6h-2v3h2V3M7 3H5v3h2V3m16 4h-3v2h3V7m-4 0H5v10h14V7M4 7H1v2h3V7m19 8h-3v2h3v-2M4 15H1v2h3v-2m15 3h-2v3h2v-3M7 18H5v3h2v-3Z"/>`),
		g.Group(children),
	)
}