package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRightThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M27.317 18.846c2.242-1.235 2.242-4.457 0-5.693L7.82 2.403C5.653 1.21 3 2.777 3 5.25v21.492c0 2.473 2.652 4.04 4.818 2.846l19.499-10.742Z"/>`),
		g.Group(children),
	)
}