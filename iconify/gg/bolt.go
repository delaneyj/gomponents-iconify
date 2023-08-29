package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 21.5l8.5-8.5l-4.5-3l2-7.5L6.5 11l4.5 3l-2 7.5Z"/>`),
		g.Group(children),
	)
}