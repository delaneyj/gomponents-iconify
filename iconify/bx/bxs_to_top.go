package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsToTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M5 15h4v6h6v-6h4l-7-8zM4 3h16v2H4z" fill="currentColor"/>`),
		g.Group(children),
	)
}