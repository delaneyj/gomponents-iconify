package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkewMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.5 11l-2.09 9H5.5l2.09-9h4.91M15 9H6L3 22h9l3-13m6-3l-4-4v3H9v2h8v3l4-4Z"/>`),
		g.Group(children),
	)
}