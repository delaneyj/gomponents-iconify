package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 19.1l2.5 1.5l-.7-2.8l2.2-1.9l-2.9-.2L17 13l-1.1 2.6l-2.9.3l2.2 1.9l-.7 2.8l2.5-1.5M3 14h8v2H3v-2m0-8h12v2H3V6m0 4h12v2H3v-2Z"/>`),
		g.Group(children),
	)
}