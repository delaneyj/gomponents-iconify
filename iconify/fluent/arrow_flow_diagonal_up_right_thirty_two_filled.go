package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFlowDiagonalUpRightThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 5a1 1 0 1 0 0 2h6.586L11.618 18.968a5.5 5.5 0 1 0 1.414 1.414L25 8.415V15a1 1 0 1 0 2 0V6a1 1 0 0 0-1-1h-9Z"/>`),
		g.Group(children),
	)
}