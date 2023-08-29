package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpNoSim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.79 3.74L2.38 5.15l2.74 2.74l-.12.12V21h13.27l1.58 1.62l1.41-1.41zM19 16.11V3h-8.99L7.95 5.06z"/>`),
		g.Group(children),
	)
}