package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.682 13.15c-2.242 1.236-2.243 4.457-.001 5.693l19.498 10.75c2.166 1.194 4.82-.373 4.82-2.846V5.255c0-2.473-2.653-4.04-4.819-2.847L4.682 13.15Z"/>`),
		g.Group(children),
	)
}