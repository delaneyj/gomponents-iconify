package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsArrowToTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M6 4h12v2H6zm5 10v6h2v-6h5l-6-6l-6 6z" fill="currentColor"/>`),
		g.Group(children),
	)
}