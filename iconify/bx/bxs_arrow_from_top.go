package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsArrowFromTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M6 4h12v2H6zm5 4v6H6l6 6l6-6h-5V8z" fill="currentColor"/>`),
		g.Group(children),
	)
}