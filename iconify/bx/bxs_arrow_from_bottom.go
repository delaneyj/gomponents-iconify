package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsArrowFromBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M6 18h12v2H6zm6-14l-6 6h5v6h2v-6h5z" fill="currentColor"/>`),
		g.Group(children),
	)
}