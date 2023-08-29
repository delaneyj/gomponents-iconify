package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsCableCar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M2 9.76l9-2.45V10H7a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2h-4V6.76l9-2.45V2.24L2 7.69zM11 12v3H7v-3zm6 0v3h-4v-3z" fill="currentColor"/>`),
		g.Group(children),
	)
}