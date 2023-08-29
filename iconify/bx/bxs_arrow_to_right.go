package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsArrowToRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M18 6h2v12h-2zm-8 5H4v2h6v5l6-6l-6-6z" fill="currentColor"/>`),
		g.Group(children),
	)
}