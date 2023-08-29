package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 13a9 9 0 0 1 18 0v3.807l1.928 4.822A1 1 0 0 1 26 23H6a1 1 0 0 1-.928-1.371L7 16.807V13Zm5.125 12A3.99 3.99 0 0 0 16 28a3.99 3.99 0 0 0 3.875-3h-7.75Z"/>`),
		g.Group(children),
	)
}