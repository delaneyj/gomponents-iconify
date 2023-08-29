package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsHAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.243 7.757l1.414 1.415L3.828 11h16.344l-1.829-1.828l1.414-1.415L24 12l-4.243 4.243l-1.414-1.415L20.171 13H3.828l1.829 1.828l-1.414 1.415L0 12l4.243-4.243Z"/>`),
		g.Group(children),
	)
}