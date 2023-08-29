package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsArrowFromLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M4 6h2v12H4zm10 5H8v2h6v5l6-6l-6-6z" fill="currentColor"/>`),
		g.Group(children),
	)
}