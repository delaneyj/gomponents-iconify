package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownThick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 4h4v9l3.5-3.5l2.42 2.42L12 19.84l-7.92-7.92L6.5 9.5L10 13V4Z"/>`),
		g.Group(children),
	)
}