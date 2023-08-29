package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCompact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19h6v-7H3v7m7 0h12v-7H10v7M3 5v6h19V5H3Z"/>`),
		g.Group(children),
	)
}