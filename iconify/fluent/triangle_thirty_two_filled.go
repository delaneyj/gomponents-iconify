package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.847 4.684c-1.235-2.242-4.457-2.243-5.693-.001L2.404 24.18C1.21 26.346 2.777 29 5.251 29h21.492c2.473 0 4.04-2.653 2.846-4.819L18.847 4.684Z"/>`),
		g.Group(children),
	)
}