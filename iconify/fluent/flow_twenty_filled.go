package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.041 7.5H12A1.5 1.5 0 0 0 10.5 9v2A2.5 2.5 0 0 1 8 13.5h-.042a3 3 0 1 1 0-1H8A1.5 1.5 0 0 0 9.5 11V9A2.5 2.5 0 0 1 12 6.5h.041a3 3 0 1 1 0 1Z"/>`),
		g.Group(children),
	)
}