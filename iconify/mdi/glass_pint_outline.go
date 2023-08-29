package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassPintOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 2l2 20h11l2-20H4m2.2 2h10.6l-1.6 16H7.8L6.2 4Z"/>`),
		g.Group(children),
	)
}