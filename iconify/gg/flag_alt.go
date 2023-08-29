package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12.439 7l2.4-3H7v6h7.839l-2.4-3ZM19 12H7v10H5V2h14l-4 5l4 5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}