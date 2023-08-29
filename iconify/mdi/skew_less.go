package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkewLess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.5 11l-2.09 9H10.5l2.09-9h4.91M20 9h-9L8 22h9l3-13M4 6l4-4v3h8v2H8v3L4 6Z"/>`),
		g.Group(children),
	)
}