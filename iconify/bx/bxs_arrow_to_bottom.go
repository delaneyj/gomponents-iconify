package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsArrowToBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M6 18h12v2H6zm5-14v6H6l6 6l6-6h-5V4z" fill="currentColor"/>`),
		g.Group(children),
	)
}