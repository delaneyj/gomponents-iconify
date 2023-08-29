package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineDivide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 11h18v2H3zm9.003-8a3 3 0 1 1-.006 6a3 3 0 0 1 .006-6zM12 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm.003 10a3 3 0 1 1-.006 6a3 3 0 0 1 .006-6zM12 17a1 1 0 1 0 0 2a1 1 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}