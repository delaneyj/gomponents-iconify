package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsArrowToLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M4 6h2v12H4zm10 7h6v-2h-6V6l-6 6l6 6z" fill="currentColor"/>`),
		g.Group(children),
	)
}