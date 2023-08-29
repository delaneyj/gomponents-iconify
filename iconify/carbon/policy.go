package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Policy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 18a6 6 0 1 0-10 4.46V30l4-1.893L28 30v-7.54A5.98 5.98 0 0 0 30 18zm-4 8.84l-2-.947l-2 .947v-3.19a5.888 5.888 0 0 0 4 0zM24 22a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zM9 14h7v2H9zm0-6h10v2H9z"/><path fill="currentColor" d="M6 30a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h16a2.002 2.002 0 0 1 2 2v4h-2V4H6v24h10v2Z"/>`),
		g.Group(children),
	)
}